package room

import (
	"guestbook/entity"
)

type DrawingRoom struct {
	canvasId entity.ID
	inbox    chan Message
	join     chan *Client
	leave    chan *Client
	clients  Clients
}

func openDrawingRoom(canvasId entity.ID, client *Client) *DrawingRoom {
	room := &DrawingRoom{
		canvasId: canvasId,
		inbox:    make(chan Message),
		join:     make(chan *Client),
		leave:    make(chan *Client),
		clients:  make(Clients),
	}

	room.clients.Insert(client)
	// TODO : run을 돌리면서 room 정보를 반환하는 방법?
	go room.run()
	return room
}

func (room *DrawingRoom) run() {
	for {
		select {
		case client := <-room.join:
			room.clients.Insert(client)
		case client := <-room.leave:
			room.clients.Remove(client)
			close(client.inbox)
			// TODO : 아직 메시지가 남아있는 경우 Room을 즉시 닫지 않을 수 있을까?
			if room.isEmpty() {
				room.closeRoom()
				return
			}
		case message := <-room.inbox:
			for client := range room.clients {
				client.inbox <- message
			}
		}
	}
}

func (room *DrawingRoom) isEmpty() bool {
	return room.clients.IsEmpty()
}

func (room *DrawingRoom) closeRoom() {
	close(room.inbox)
	close(room.join)
	close(room.leave)
}
