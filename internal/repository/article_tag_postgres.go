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

func (a ArticleTagRepo) GetTagIdsByArticleId(ctx context.Context, articleId uint64) ([]uint64, error) {
	rows, err := a.Conn.Query(ctx, "SELECT tag_id FROM articles_tags WHERE article_id = $1", articleId)
	if err != nil {
		return nil, err
	}
	var tagIds []uint64
	for rows.Next() {
		var tagId uint64
		if err := rows.Scan(&tagId); err != nil {
			return nil, err
		}
		tagIds = append(tagIds, tagId)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return tagIds, nil
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
