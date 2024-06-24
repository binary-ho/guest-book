package main

import (
	"guestbook/entity/user"
	"net/http"
)

func (app *application) routes() {
	controllers := app.controllers
	serveMux := http.NewServeMux()

	serveUserApis(controllers.userController, serveMux)
}

func serveUserApis(controller *user.Controller, serveMux *http.ServeMux) {
	serveMux.HandleFunc(controller.SignUp())
}
