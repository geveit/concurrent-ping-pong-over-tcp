package main

import (
	"fmt"
	"net"
)

type client struct {
	hub  *hub
	conn net.Conn
	send chan []byte
}

func newClient(hub *hub, conn net.Conn) *client {
	return &client{
		hub:  hub,
		conn: conn,
		send: make(chan []byte, 256),
	}
}

func (c *client) readFrom() {
	defer func() {
		fmt.Println("Connection closed for client", c.conn.RemoteAddr())
		c.hub.unregister <- c
		c.conn.Close()
	}()

	buffer := make([]byte, 4)
	for {
		_, err := c.conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from client", c.conn.RemoteAddr(), ":", err)
			return
		}
		fmt.Printf("Received %s from %s\n", string(buffer), c.conn.RemoteAddr())
	}
}

func (c *client) writeTo() {
	defer func() {
		c.conn.Close()
		fmt.Println("Connection closed for client", c.conn.RemoteAddr())
	}()

	for message := range c.send {
		_, err := c.conn.Write(message)
		if err != nil {
			fmt.Println("Error writing to client", c.conn.RemoteAddr(), ":", err)
			return
		}
	}
}
