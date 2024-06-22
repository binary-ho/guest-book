package room

import (
	"guestbook/entity"
)

// TODO : 추후 Redis로 관리
type drawingRooms struct {
	rooms map[entity.ID]*DrawingRoom
}

var roomsInstance drawingRooms

func init() {
	roomsInstance.rooms = make(map[entity.ID]*DrawingRoom)
}

func OpenDrawingRoom(id entity.ID, client *Client) {
	roomsInstance.rooms[id] = openDrawingRoom(id, client)
}

func GetDrawingRoom(id entity.ID) *DrawingRoom {
	return roomsInstance.rooms[id]
}

func CloseDrawingRoom(id entity.ID) {
	delete(roomsInstance.rooms, id)
}
