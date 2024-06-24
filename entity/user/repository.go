package user

import (
	"database/sql"
	"errors"
	"guestbook/entity"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

const (
	SelectUserById       = `SELECT id, handle, nickname, githubId, profileUrl, plan FROM users WHERE id = ?`
	SelectUserByGithubId = `SELECT id, handle, nickname, githubId, profileUrl, plan FROM users WHERE githubId = ?`
	InsertUser           = `INSERT INTO users (id, handle, nickname, githubId, profileUrl, plan) VALUES (?, ?, ?, ?, ?, ?)`
	UpdateUser           = `UPDATE users SET handle = ?, nickname = ?, githubId = ?, profileUrl = ?, plan = ? WHERE id = ?`
)

func (repo *Repository) FindById(id entity.ID) (*Entity, error) {
	row := repo.db.QueryRow(SelectUserById, int64(id))
	return scanRow(row)
}

func (repo *Repository) ExistsByGithubId(githubId entity.ID) bool {
	row := repo.db.QueryRow(SelectUserByGithubId, int64(githubId))
	_, err := scanRow(row)
	return err == nil || !errors.Is(err, sql.ErrNoRows)
}

func (repo *Repository) Save(user Entity) (*Entity, error) {
	if user.ID() == entity.DefaultId() {
		return insertUser(user, repo)
	}
	_, err := repo.db.Exec(UpdateUser, user.handle, user.nickname, user.githubId, user.profileUrl, user.plan, user.id)
	return &user, err
}

func insertUser(user Entity, repo *Repository) (*Entity, error) {
	result, err := repo.db.Exec(
		InsertUser, user.id, user.handle, user.nickname, user.githubId, user.profileUrl, user.plan)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	return &Entity{
		id:         entity.ID(id),
		handle:     user.handle,
		nickname:   user.nickname,
		githubId:   user.githubId,
		profileUrl: user.profileUrl,
		plan:       user.Plan(),
	}, nil
}

func scanRow(row *sql.Row) (*Entity, error) {
	var user Entity
	return &user, row.Scan(&user.id, &user.handle, &user.nickname, &user.githubId, &user.profileUrl, &user.plan)
}
