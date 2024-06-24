package main

import "guestbook/entity/user"

type controllers struct {
	userController *user.Controller
}

func (app *application) initControllers() {
	services := app.services
	app.controllers = &controllers{
		userController: user.InitController(services.userService),
	}
}
