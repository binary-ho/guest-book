package entities

import "guestbook/entity"

type User struct {
	id entity.Id
}

func (user *User) Id() entity.Id {
	return user.id
}
