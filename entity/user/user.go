package user

import (
	"guestbook/entity"
	"net/url"
)

type Entity struct {
	id         entity.ID
	handle     string
	nickname   string
	profileUrl url.URL
	plan       plan
}

func (user *Entity) ID() entity.ID {
	return user.id
}

func (user *Entity) Plan() plan {
	return user.plan
}
