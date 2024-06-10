package register

import (
	"guestbook/entity"
)

type Canvas struct {
	id   entity.ID
	data map[uint32]LWWRegister
}
