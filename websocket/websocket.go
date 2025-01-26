package websocket

import (
	"sync"

	"github.com/gofiber/websocket/v2"
)

var clients = make(map[*websocket.Conn]bool)
var mutex = sync.Mutex{}

func WebSocketHandler(c *websocket.Conn) {
	mutex.Lock()
	clients[c] = true
	mutex.Unlock()

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			mutex.Lock()
			delete(clients, c)
			mutex.Unlock()
			c.Close()
			break
		}

		broadcastMessage(msg)
	}
}

func broadcastMessage(message []byte) {
	mutex.Lock()
	defer mutex.Unlock()

	for client := range clients {
		client.WriteMessage(1, message)
	}
}
