package component

import (
	"guestbook/entity/user"
)

type Services struct {
	UserService *user.Service
}

func InitServices(repository *user.Repository) *Services {
	return &Services{
		UserService: user.InitService(repository),
	}
}
