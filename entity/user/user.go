package entities

import "guestbook/entity"

type User struct {
	id   entity.Id
	plan plan
}

func (user *User) Id() entity.Id {
	return user.id
}

func (user *User) Plan() plan {
	return user.plan
}
