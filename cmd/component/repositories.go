package component

import (
	"database/sql"
	"guestbook/entity/user"
)

type Repositories struct {
	UserRepository *user.Repository
}

func InitRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		UserRepository: user.NewRepository(db),
	}
}
