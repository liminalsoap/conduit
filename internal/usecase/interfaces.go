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
	ArticleTag
	Like
}

type Tag interface {
	List(context.Context) (*[]entity.Tag, error)
	GetByTitle(context.Context, string) (uint64, error)
	GetByTitles(context.Context, []string) ([]uint64, error)
	GetByIds(context.Context, []uint64) ([]string, error)
}

type TagRepo interface {
	GetTags(context.Context) (*[]entity.Tag, error)
	GetByTitle(context.Context, string) (uint64, error)
	GetByTitles(context.Context, []string) ([]uint64, error)
	GetByIds(context.Context, []uint64) ([]string, error)
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
	GetBySlug(context.Context, string) (entity.ArticleInput, error)
	Create(context.Context, entity.Article, []string) (entity.ArticleInput, error)
	Update(context.Context, entity.Article, string) error
	DeleteBySlug(context.Context, string) error
	List(context.Context) ([]entity.ArticleInput, error)

	GetTagList(context.Context, uint64) ([]string, error)
}

type ArticleRepo interface {
	GetIdBySlug(context.Context, string) (uint64, error)
	GetBySlug(context.Context, string) (entity.ArticleInput, error)
	Create(context.Context, entity.Article) (uint64, error)
	Update(context.Context, entity.Article, string) error
	DeleteBySlug(context.Context, string) error

	List(context.Context) ([]entity.ArticleInput, error)
}

type ArticleTag interface {
	GetTagIdsByArticleId(context.Context, uint64) ([]uint64, error)
	Add(context.Context, uint64, uint64) error
	AddList(context.Context, uint64, []uint64) error
}

type ArticleTagRepo interface {
	GetTagIdsByArticleId(context.Context, uint64) ([]uint64, error)
	Add(context.Context, uint64, uint64) error
}

type Like interface {
	Favorite(context.Context, uint64, uint64) error
	Unfavorite(context.Context, uint64, uint64) error
	Count(context.Context, uint64) (uint64, error)
}

type LikeRepo interface {
	Add(context.Context, uint64, uint64) error
	Delete(context.Context, uint64, uint64) error
	Count(context.Context, uint64) (uint64, error)
}
