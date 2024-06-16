package room

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	room       *DrawingRoom
	connection *websocket.Conn
	inbox      chan Message
}

func (client *Client) read() {
	defer client.connection.Close()
	for {
		var message Message
		err := client.connection.ReadJSON(&message)
		if err != nil {
			return
		}
		client.room.inbox <- message
	}
}

func (client *Client) write() {
	defer client.connection.Close()
	var message Message
	err := client.connection.WriteJSON(&message)
	if err != nil {
		return
	}
}
