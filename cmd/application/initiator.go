package application

import (
	"database/sql"
	"flag"
	"guestbook/cmd/component"
	"guestbook/cmd/config"
	"log/slog"
	"os"
)

func InitApplication() *Application {
	configuration := config.ApplicationConfiguration()
	application := &Application{
		logger: getNewLogger(),
		config: configuration,
		db:     getDbConnectionPool(configuration.Db.DataSource(), configuration.Db.Driver()),
	}

	application.repositories = component.InitRepositories(application.db)
	application.services = component.InitServices(application.repositories.UserRepository)
	application.controllers = component.InitControllers(application.services.UserService)
	return application
}

func getNewLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, nil))
}

func getDbConnectionPool(dataSource, driver string) *sql.DB {
	dataSourceFlag := flag.String("dataSourceFormat", dataSource, "Data Source Name")
	flag.Parse()

	db, err := sql.Open(driver, *dataSourceFlag)
	if err != nil {
		panic(err)
	}
	healthCheck(db)
	return db
}

func healthCheck(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		db.Close()
		panic(err)
	}
}
