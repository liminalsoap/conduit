package repository

import (
	"conduit-go/internal/entity"
	"conduit-go/pkg/postgres"
	"context"
)

type UserRepo struct {
	*postgres.Postgres
}

func NewUserRepo(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

func (r UserRepo) Create(ctx context.Context, user entity.User) (entity.User, error) {
	_, err := r.Conn.Exec(
		ctx,
		"INSERT INTO users(email, username, password) VALUES($1, $2, $3)",
		user.Email,
		user.Username,
		user.Password,
	)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r UserRepo) GetById(ctx context.Context, id uint64) (entity.User, error) {
	var user entity.User
	err := r.Conn.QueryRow(
		ctx,
		"SELECT id, email, username, password, bio, image FROM users WHERE id = $1",
		id,
	).Scan(
		&user.Id,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.Bio,
		&user.Image,
	)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r UserRepo) Update(ctx context.Context, user entity.User) (entity.User, error) {
	_, err := r.Conn.Exec(ctx,
		"UPDATE users SET email = $1, username = $2, password = $3, image = $4, bio = $5 WHERE id = $6",
		user.Email,
		user.Username,
		user.Password,
		user.Image,
		user.Bio,
		user.Id,
	)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r UserRepo) FindByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	err := r.Conn.QueryRow(
		ctx,
		"SELECT id, email, username, password, bio, image FROM users WHERE email = $1",
		email,
	).Scan(
		&user.Id,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.Bio,
		&user.Image,
	)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r UserRepo) FindByUsername(ctx context.Context, username string) (entity.User, error) {
	var user entity.User
	err := r.Conn.QueryRow(
		ctx,
		"SELECT id, email, username, password, bio, image FROM users WHERE username = $1",
		username,
	).Scan(
		&user.Id,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.Bio,
		&user.Image,
	)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
