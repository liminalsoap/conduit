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

func (a ArticleRepo) GetIdBySlug(ctx context.Context, slug string) (uint64, error) {
	var id uint64
	err := a.Conn.QueryRow(ctx, "SELECT id FROM articles WHERE slug = $1", slug).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (a ArticleRepo) GetBySlug(ctx context.Context, slug string) (entity.Article, error) {
	var article entity.Article
	err := a.Conn.QueryRow(
		ctx,
		"SELECT id, slug, title, description, body, created_at, updated_at, user_id FROM articles WHERE slug = $1",
		slug,
	).Scan(
		&article.Id,
		&article.Slug,
		&article.Title,
		&article.Description,
		&article.Body,
		&article.CreatedAt,
		&article.UpdatedAt,
		&article.UserId,
	)
	if err != nil {
		return entity.Article{}, err
	}
	return article, nil
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
	createdArticle, err := a.GetBySlug(ctx, article.Slug)
	if err != nil {
		return entity.Article{}, err
	}
	return createdArticle, err
}

func (a ArticleRepo) Update(ctx context.Context, article entity.Article, slug string) error {
	_, err := a.Conn.Exec(
		ctx,
		"UPDATE articles SET title = $1, description = $2, body = $3 WHERE id = $4",
		article.Title,
		article.Description,
		article.Body,
		article.Id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (a ArticleRepo) DeleteBySlug(ctx context.Context, slug string) error {
	_, err := a.Conn.Exec(ctx, "DELETE FROM articles WHERE slug = $1", slug)
	if err != nil {
		return err
	}
	return nil
}

func (a ArticleRepo) List(ctx context.Context, slug string) ([]entity.Article, error) {
	//TODO implement me
	panic("implement me")
}
