package main

import (
	"fmt"
	"net"
)

func main() {
	h := newHub()
	go h.run()
	go h.broadcastPong()

	listener, err := net.Listen("tcp", ":5555")
	if err != nil {
		fmt.Println("Error creating listener:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on port 5555")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			return
		}
		h.registerClient(conn)
	}
}
