package usecase

import (
	"conduit-go/internal/entity"
	"context"
)

type UseCases struct {
	Tag
	User
	Following
	Article
}

type Tag interface {
	List(context.Context) (*[]entity.Tag, error)
	GetByTitle(context.Context, string) (uint64, error)
	GetByTitles(context.Context, []string) ([]uint64, error)
}

type TagRepo interface {
	GetTags(context.Context) (*[]entity.Tag, error)
	GetByTitle(context.Context, string) (uint64, error)
	GetByTitles(context.Context, []string) ([]uint64, error)
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

type Article interface {
	GetBySlug(context.Context, string) (entity.Article, error)
	Create(context.Context, entity.Article, []string) (entity.Article, error)
	Update(context.Context, string) (entity.Article, error)
	DeleteBySlug(context.Context, string) (entity.Article, error)
	List(context.Context, string) ([]entity.Article, error)
}

type ArticleRepo interface {
	GetBySlug(context.Context, string) (entity.Article, error)
	Create(context.Context, entity.Article) (entity.Article, error)
	Update(context.Context, string) (entity.Article, error)
	DeleteBySlug(context.Context, string) (entity.Article, error)
	List(context.Context, string) ([]entity.Article, error)
}

type ArticleTag interface {
	Add(context.Context, uint64, uint64) error
	AddList(context.Context, uint64, []uint64) error
}

type ArticleTagRepo interface {
	Add(context.Context, uint64, uint64) error
}
