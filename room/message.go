package room

import (
	"guestbook/entity/user"
	"guestbook/register"
)

type Message struct {
	user      user.Entity
	timestamp register.Timestamp
	pixels    []PixelData
}
