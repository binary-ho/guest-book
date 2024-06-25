package component

import (
	"guestbook/entity/user"
)

type Controllers struct {
	UserController *user.Controller
}

func InitControllers(userService *user.Service) *Controllers {
	return &Controllers{
		UserController: user.InitController(userService),
	}
}
