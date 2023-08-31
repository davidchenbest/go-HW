package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Client connected")

	r := bufio.NewReader(conn)

	// buffer := make([]byte, 1024)

	for {
		// n, err := conn.Read(buffer)
		n, err := r.ReadString('\n')

		if err != nil {

			fmt.Println("Error reading:", err)
			break
		}

		// message := string(buffer[:n])
		message := string(n)

		message = strings.TrimSpace(message)

		fmt.Printf("Received: %s\n", message)

		// Simulate WebSocket-like response
		response := fmt.Sprintf("You sent: %s\n", message)
		_, err = conn.Write([]byte(response))
		conn.Write([]byte("k"))
		if err != nil {
			fmt.Println("Error writing:", err)
			break
		}
	}

	fmt.Println("Client disconnected")
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("TCP server started on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}
