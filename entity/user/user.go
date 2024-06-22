package user

import (
	"guestbook/common"
	"guestbook/entity"
)

type Entity struct {
	id         entity.ID
	handle     string
	nickname   string
	githubId   entity.ID
	profileUrl common.Url
	plan       plan
}

func Create() *Entity {
	parsedUrl, _ := common.CreateUrl("https://www.github.com/binary-ho")
	return &Entity{
		handle:     "binary-ho",
		nickname:   "binary-ho",
		githubId:   entity.ID(77),
		profileUrl: parsedUrl,
		plan:       FreePlan(),
	}
}

func (user *Entity) ID() entity.ID {
	return user.id
}

func (user *Entity) Plan() plan {
	return user.plan
}
