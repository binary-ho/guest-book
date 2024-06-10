package user

import "guestbook/entity"

type User struct {
	id   entity.ID
	plan plan
}

func (user *User) ID() entity.ID {
	return user.id
}

func (user *User) Plan() plan {
	return user.plan
}
