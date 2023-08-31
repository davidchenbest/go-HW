package main

import (
	"fmt"
	"net"
	"strings"
)

const (
	endOfLine = "\r\n"
)

func handleWebSocket(conn net.Conn) {
	defer conn.Close()

	fmt.Println("WebSocket client connected")

	// Send WebSocket handshake response
	_, err := conn.Write([]byte("HTTP/1.1 101 Switching Protocols" + endOfLine))
	if err != nil {
		fmt.Println("Error sending response:", err)
		return
	}

	_, err = conn.Write([]byte("Upgrade: websocket" + endOfLine))
	if err != nil {
		fmt.Println("Error sending response:", err)
		return
	}

	_, err = conn.Write([]byte("Connection: Upgrade" + endOfLine))
	if err != nil {
		fmt.Println("Error sending response:", err)
		return
	}

	_, err = conn.Write([]byte(endOfLine))
	if err != nil {
		fmt.Println("Error sending response:", err)
		return
	}

	// Read WebSocket frames
	for {
		frame := make([]byte, 1024)
		n, err := conn.Read(frame)
		if err != nil {
			fmt.Println("Error reading frame:", err)
			break
		}

		message := string(frame[:n])
		message = strings.TrimSpace(message)

		fmt.Printf("Received: %s\n", message)

		// Echo the received message back as a WebSocket frame
		response := []byte{0x81, byte(len(message))}
		response = append(response, []byte(message)...)
		_, err = conn.Write(response)
		if err != nil {
			fmt.Println("Error writing response:", err)
			break
		}
	}

	fmt.Println("WebSocket client disconnected")
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("WebSocket server started on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleWebSocket(conn)
	}
}
