package main

import (
	"fmt"
	"math/rand"
	"net"
	"sync"
	"time"
)

func main() {
	const numPings = 1000
	var wg sync.WaitGroup
	wg.Add(numPings)

	for i := 0; i < numPings; i++ {
		go func(pingID int) {
			defer wg.Done()
			runPing(pingID)
		}(i)
	}

	wg.Wait()
	fmt.Println("All clients have finished.")
}

func runPing(pingID int) {
	conn, err := net.Dial("tcp", "localhost:5555")
	if err != nil {
		fmt.Printf("Client %d: Error connecting to Pong Server: %v\n", pingID, err)
		return
	}
	defer conn.Close()

	fmt.Printf("Client %d: Connected to Pong Server\n", pingID)

	go writeMessageRoutine(conn, pingID)

	readMessageRoutine(conn, pingID)
}

func writeMessageRoutine(conn net.Conn, pingID int) {
	message := []byte("ping")

	for {
		interval := rand.Float32() + 0.5
		time.Sleep(time.Duration(interval) * time.Second)

		_, err := conn.Write(message)
		if err != nil {
			fmt.Printf("Client %d: Error writing to connection: %v\n", pingID, err)
			return
		}
		fmt.Printf("Client %d: Sent ping\n", pingID)
	}
}

func readMessageRoutine(conn net.Conn, pingID int) {
	defer func() {
		conn.Close()
		fmt.Printf("Client %d: Connection closed\n", pingID)
	}()

	buffer := make([]byte, 4)

	for {
		_, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("Client %d: Error reading from connection: %v\n", pingID, err)
			return
		}
		fmt.Printf("Client %d: Received %s\n", pingID, string(buffer))
	}
}
