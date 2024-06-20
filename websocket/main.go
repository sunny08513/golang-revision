package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// upgrader is used to upgrade a HTTP connection to a WebSocket connection.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024, // Read buffer size
	WriteBufferSize: 1024, // Write buffer size
	CheckOrigin: func(r *http.Request) bool {
		// Allow connections from any origin (for simplicity)
		return true
	},
}

// handleConnections handles incoming WebSocket connections.
func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close() // Ensure the connection is closed when the function exits

	// Goroutine to send a ping message every 30 seconds to keep the connection alive
	go func() {
		for {
			// Send a ping message
			err := ws.WriteMessage(websocket.PingMessage, []byte{})
			if err != nil {
				fmt.Println("Ping error:", err)
				return
			}
			// Wait for 30 seconds before sending the next ping
			time.Sleep(30 * time.Second)
		}
	}()

	// Main loop to handle incoming messages
	for {
		// Read a message from the WebSocket connection
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("Read error:", err)
			break
		}
		fmt.Printf("Received: %s\n", message)

		// Echo the message back to the client
		if err := ws.WriteMessage(messageType, message); err != nil {
			fmt.Println("Write error:", err)
			break
		}
	}
}

// main is the entry point of the program
func main() {
	// Define the WebSocket endpoint
	http.HandleFunc("/ws", handleConnections)

	// Start the server on port 8080
	fmt.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
