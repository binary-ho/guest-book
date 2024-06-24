package main

import (
	"guestbook/entity/user"
)

type Repositories struct {
	userRepository *user.Repository
}

func (app *application) initRepositories() {
	app.repositories = &Repositories{
		userRepository: user.NewRepository(app.db),
	}
}
