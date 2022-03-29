package socket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Connection struct {
	ws *websocket.Conn
}

type ChatMessage struct {
	Message    string `json:"message"`
	ReceiverId string `json:"receiver_id"`
}

/**
message structure

<roomId1>;<roomId2>=<MESSAGE_STRING>

*/

func (c *Connection) readMessages(socket *Socket) {
	for {
		var message ChatMessage
		err := c.ws.ReadJSON(&message)
		if err != nil {
			log.Println("Error during message reading:", err)
			break
		}
		log.Printf("Received: %s", message)

		// Send to respective users
		var connections = socket.GetConnectionFromId(message.ReceiverId)

		for _, connection := range connections {
			connection.writeMessage(message)
		}

		if err != nil {
			log.Println("Error during message writing:", err)
			break
		}
	}
}

func (c *Connection) writeMessage(message ChatMessage) {
	err := c.ws.WriteJSON(message)

	if err != nil {
		fmt.Println("Some error occured while sending message - ", err)
	}
}
