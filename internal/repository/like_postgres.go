package repository

import (
	"conduit-go/pkg/postgres"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
)

type LikeRepo struct {
	*postgres.Postgres
}

func NewLikeRepo(pg *postgres.Postgres) *LikeRepo {
	return &LikeRepo{pg}
}

func (l LikeRepo) Add(ctx context.Context, articleId, userId uint64) error {
	var id uint64
	err := l.Conn.QueryRow(
		ctx,
		"SELECT id FROM likes WHERE user_id = $1 AND article_id = $2",
		userId,
		articleId,
	).Scan(&id)

	if err != nil && err != pgx.ErrNoRows {
		return err
	}
	if id != 0 {
		return errors.New("user already favorited to article")
	}

	_, err = l.Conn.Exec(ctx, "INSERT INTO likes(article_id, user_id) VALUES ($1, $2)", articleId, userId)
	if err != nil {
		return err
	}

	return nil
}

func (l LikeRepo) Delete(ctx context.Context, articleId, userId uint64) error {
	_, err := l.Conn.Exec(ctx, "DELETE FROM likes WHERE article_id = $1 AND user_id = $2", articleId, userId)
	if err != nil {
		return err
	}

	return nil
}

func (l LikeRepo) Count(ctx context.Context, articleId uint64) (uint64, error) {
	var count uint64
	err := l.Conn.QueryRow(
		ctx,
		"SELECT COUNT(*) FROM likes GROUP BY article_id HAVING article_id = $1",
		articleId,
	).Scan(&count)
	if err != nil && err != pgx.ErrNoRows {
		return 0, err
	}

	return count, nil
}
