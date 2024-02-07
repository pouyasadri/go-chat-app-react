package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

// Message define a message object
type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

// Define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			return
		}
		message := Message{Type: messageType, Body: string(p)}
		c.Pool.Broadcast <- message
		fmt.Printf("Message Received: %+v\n", message)
	}
}
