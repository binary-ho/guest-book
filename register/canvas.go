package register

import (
	"guestbook/entity"
	"guestbook/entity/user"
)

type Canvas struct {
	id    entity.ID
	owner user.User
	data  map[uint32]LWWRegister
}
