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
	//TODO implement me
	panic("implement me!")
}

func (r UserRepo) Update(ctx context.Context, user entity.User) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}
