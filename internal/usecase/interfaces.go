package usecase

import (
	"conduit-go/internal/entity"
	"context"
)

type UseCases struct {
	Tag
	User
}

type Tag interface {
	List(context.Context) (*[]entity.Tag, error)
}

type TagRepo interface {
	GetTags(context.Context) (*[]entity.Tag, error)
}

type User interface {
	Create(context.Context, entity.User) (entity.User, error)
	GetUser(context.Context, uint64) (entity.User, error)
	Update(context.Context, entity.User) (entity.User, error)
	FindByEmail(context.Context, string) (entity.User, error)
}

type UserRepo interface {
	Create(context.Context, entity.User) (entity.User, error)
	GetById(context.Context, uint64) (entity.User, error)
	Update(context.Context, entity.User) (entity.User, error)
	FindByEmail(context.Context, string) (entity.User, error)
}
