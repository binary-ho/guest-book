package application

import (
	"guestbook/entity/user"
	"net/http"
)

func (app *Application) routes() *http.ServeMux {
	controllers := app.controllers
	serveMux := http.NewServeMux()

	serveUserApis(controllers.UserController, serveMux)
	return serveMux
}

func serveUserApis(controller *user.Controller, serveMux *http.ServeMux) {
	serveMux.HandleFunc(controller.SignUp())
}
