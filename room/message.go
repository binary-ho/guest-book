package room

import (
	"guestbook/entity/user"
	"guestbook/register"
)

type Message struct {
	user      user.User
	timestamp register.Timestamp
	pixels    []PixelData
}
