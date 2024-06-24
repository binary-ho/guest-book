package main

import (
	"database/sql"
	"log/slog"
)

type application struct {
	logger *slog.Logger
	config Config
	db     *sql.DB

	controllers  *controllers
	services     *Services
	repositories *Repositories
}

func InitApplicationContext(logger *slog.Logger, config Config, db *sql.DB) *application {
	app := &application{
		logger: logger,
		config: config,
		db:     db,
	}

	app.initRepositories()
	app.initServices()
	app.initControllers()
	return app
}
