package repository

import (
	"conduit-go/pkg/postgres"
	"context"
)

type ArticleTagRepo struct {
	*postgres.Postgres
}

func NewArticleTagRepo(pg *postgres.Postgres) *ArticleTagRepo {
	return &ArticleTagRepo{pg}
}

func (a ArticleTagRepo) Add(ctx context.Context, articleId uint64, tagId uint64) error {
	_, err := a.Conn.Exec(
		ctx,
		"INSERT INTO articles_tags(article_id, tag_id) VALUES ($1, $2)",
		articleId,
		tagId,
	)
	if err != nil {
		return err
	}
	return nil
}
