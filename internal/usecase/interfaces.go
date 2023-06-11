package usecase

import (
	"conduit-go/internal/entity"
	"context"
)

type UseCases struct {
	Tag
	User
	Following
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
	FindByUsername(context.Context, string) (entity.User, error)
}

type UserRepo interface {
	Create(context.Context, entity.User) (entity.User, error)
	GetById(context.Context, uint64) (entity.User, error)
	Update(context.Context, entity.User) (entity.User, error)
	FindByEmail(context.Context, string) (entity.User, error)
	FindByUsername(context.Context, string) (entity.User, error)
}

type Following interface {
	Follow(context.Context, uint64, uint64) error
	Unfollow(context.Context, uint64, uint64) error
	CheckIsFollowing(context.Context, uint64, uint64) (bool, error)
}

type FollowingRepo interface {
	Follow(context.Context, uint64, uint64) error
	Unfollow(context.Context, uint64, uint64) error
	CheckIsFollowing(context.Context, uint64, uint64) (bool, error)
}
