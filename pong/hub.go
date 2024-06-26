package main

import (
	"fmt"
	"math/rand/v2"
	"net"
	"sync"
	"time"
)

type hub struct {
	sync.RWMutex

	clients    map[*client]bool
	register   chan *client
	unregister chan *client
	broadcast  chan []byte
}

func newHub() *hub {
	return &hub{
		clients:    make(map[*client]bool),
		register:   make(chan *client),
		unregister: make(chan *client),
		broadcast:  make(chan []byte, 4),
	}
}

func (h *hub) registerClient(conn net.Conn) {
	client := newClient(h, conn)
	h.register <- client
}

func (h *hub) broadcastPong() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in hub broadcast:", r)
			panic(r)
		}
	}()
	for {
		pong := []byte("pong")
		interval := rand.Float32() + 0.5
		time.Sleep(time.Duration(interval) * time.Second)
		h.broadcast <- pong
	}
}

func (h *hub) run() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in hub run:", r)
			panic(r)
		}
	}()
	for {
		select {
		case client := <-h.register:
			h.Lock()
			h.clients[client] = true
			h.Unlock()
			fmt.Println("New client registered:", client.conn.RemoteAddr())
			go client.readFrom()
			go client.writeTo()

		case client := <-h.unregister:
			h.Lock()
			if _, ok := h.clients[client]; ok {
				close(client.send)
				delete(h.clients, client)
				fmt.Println("Client unregistered:", client.conn.RemoteAddr())
			}
			h.Unlock()

		case message := <-h.broadcast:
			h.RLock()
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
			h.RUnlock()
		}
	}
}
