package repository

import (
	"conduit-go/internal/entity"
	"conduit-go/pkg/postgres"
	"context"
)

type CommentRepo struct {
	*postgres.Postgres
}

func NewCommentRepo(pg *postgres.Postgres) *CommentRepo {
	return &CommentRepo{pg}
}

func (c CommentRepo) GetLastComment(ctx context.Context) (entity.Comment, error) {
	var comment entity.Comment
	err := c.Conn.QueryRow(
		ctx,
		"SELECT id, created_at, updated_at, body, user_id, article_id FROM comments ORDER BY created_at LIMIT 1",
	).Scan(
		&comment.Id,
		&comment.CreatedAt,
		&comment.UpdatedAt,
		&comment.Body,
		&comment.UserId,
		&comment.ArticleId,
	)
	if err != nil {
		return entity.Comment{}, err
	}
	return comment, nil
}

func (c CommentRepo) Add(ctx context.Context, articleId uint64, comment entity.Comment) (entity.Comment, error) {
	_, err := c.Conn.Exec(
		ctx,
		"INSERT INTO comments(body, user_id, article_id) VALUES($1, $2, $3)",
		comment.Body,
		comment.UserId,
		articleId,
	)
	if err != nil {
		return entity.Comment{}, err
	}
	createdComment, err := c.GetLastComment(ctx)
	if err != nil {
		return entity.Comment{}, err
	}
	return createdComment, nil
}

func (c CommentRepo) GetByArticleId(ctx context.Context, articleId uint64) ([]entity.CommentInput, error) {
	rows, err := c.Conn.Query(
		ctx,
		`SELECT c.id, c.created_at, c.updated_at, c.body, u.username, u.bio, u.image
FROM comments c
JOIN users u ON u.id = c.user_id
WHERE article_id = $1`,
		articleId,
	)
	if err != nil {
		return nil, err
	}

	var comments []entity.CommentInput
	for rows.Next() {
		var comment entity.CommentInput
		err = rows.Scan(
			&comment.Id,
			&comment.CreatedAt,
			&comment.UpdatedAt,
			&comment.Body,
			&comment.Author.Username,
			&comment.Author.Bio,
			&comment.Author.Image,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	if rows.Err() != nil {
		return nil, err
	}

	return comments, nil
}

func (c CommentRepo) Delete(ctx context.Context, id uint64) error {
	_, err := c.Conn.Exec(ctx, "DELETE FROM comments WHERE article_id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (c CommentRepo) GetById(ctx context.Context, id uint64) (entity.Comment, error) {
	var comment entity.Comment
	err := c.Conn.QueryRow(
		ctx,
		"SELECT id, created_at, updated_at, body, user_id, article_id FROM comments WHERE id = $1",
		id,
	).Scan(
		&comment.Id,
		&comment.CreatedAt,
		&comment.UpdatedAt,
		&comment.Body,
		&comment.UserId,
		&comment.ArticleId,
	)
	if err != nil {
		return entity.Comment{}, err
	}
	return comment, nil
}
