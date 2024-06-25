package application

import (
	"database/sql"
	"guestbook/cmd/component"
	"guestbook/cmd/config"
	"log/slog"
	"net/http"
)

type Application struct {
	logger *slog.Logger
	config *config.ApplicationConfig
	db     *sql.DB

	controllers  *component.Controllers
	services     *component.Services
	repositories *component.Repositories
}

func (app *Application) Run() {
	serverConfig := app.config.ServerConfig
	err := http.ListenAndServe(serverConfig.ServerAddress(), app.routes())
	if err != nil {
		app.logger.Error(err.Error())
	}
}
