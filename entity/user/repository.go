package user

import (
	"database/sql"
	"errors"
	"guestbook/entity"
)

type Repository struct {
	DB *sql.DB
}

const (
	SELECT_USER_BY_ID        = `SELECT id, handle, nickname, githubId, profileUrl, plan FROM users WHERE id = ?`
	SELECT_USER_BY_GITHUB_ID = `SELECT id, handle, nickname, githubId, profileUrl, plan FROM users WHERE githubId = ?`
	INSERT_USER              = `INSERT INTO users (id, handle, nickname, githubId, profileUrl, plan) VALUES (?, ?, ?, ?, ?, ?)`
	UPDATE_USER              = `UPDATE users SET handle = ?, nickname = ?, githubId = ?, profileUrl = ?, plan = ? WHERE id = ?`
)

func (repo *Repository) FindById(id entity.ID) (*Entity, error) {
	row := repo.DB.QueryRow(SELECT_USER_BY_ID, int64(id))
	return scanRow(row)
}

func (repo *Repository) ExistsByGithubId(githubId entity.ID) bool {
	row := repo.DB.QueryRow(SELECT_USER_BY_GITHUB_ID, int64(githubId))
	_, err := scanRow(row)
	return err == nil || !errors.Is(err, sql.ErrNoRows)
}

func (repo *Repository) Save(user Entity) (*Entity, error) {
	if user.ID() == entity.DefaultId() {
		return insertUser(user, repo)
	}
	_, err := repo.DB.Exec(UPDATE_USER, user.handle, user.nickname, user.githubId, user.profileUrl, user.plan, user.id)
	return &user, err
}

func insertUser(user Entity, repo *Repository) (*Entity, error) {
	result, err := repo.DB.Exec(
		INSERT_USER, user.id, user.handle, user.nickname, user.githubId, user.profileUrl, user.plan)
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
