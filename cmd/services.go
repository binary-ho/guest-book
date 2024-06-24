package main

import "guestbook/entity/user"

type Services struct {
	userService *user.Service
}

func (app *application) initServices() {
	repositories := app.repositories
	app.services = &Services{
		userService: user.InitService(repositories.userRepository),
	}
}
