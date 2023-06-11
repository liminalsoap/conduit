package repository

import (
	"conduit-go/internal/entity"
	"conduit-go/pkg/postgres"
	"context"
)

type ArticleRepo struct {
	*postgres.Postgres
}

func NewArticleRepo(pg *postgres.Postgres) *ArticleRepo {
	return &ArticleRepo{pg}
}

func (a ArticleRepo) GetBySlug(ctx context.Context, s string) (entity.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleRepo) Create(ctx context.Context, article entity.Article) (entity.Article, error) {
	_, err := a.Conn.Exec(
		ctx,
		"INSERT INTO articles(slug, title, description, body, user_id) VALUES ($1, $2, $3, $4, $5)",
		article.Slug,
		article.Title,
		article.Description,
		article.Body,
		article.UserId,
	)
	if err != nil {
		return entity.Article{}, err
	}
	return article, err
}

func (a ArticleRepo) Update(ctx context.Context, s string) (entity.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleRepo) DeleteBySlug(ctx context.Context, s string) (entity.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleRepo) List(ctx context.Context, s string) ([]entity.Article, error) {
	//TODO implement me
	panic("implement me")
}
