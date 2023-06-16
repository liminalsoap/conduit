package repository

import (
	"conduit-go/internal/entity"
	"conduit-go/pkg/postgres"
	"context"
	"github.com/jackc/pgx/v5"
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

func (a ArticleRepo) GetBySlug(ctx context.Context, slug string) (entity.ArticleInput, error) {
	var article entity.ArticleInput
	err := a.Conn.QueryRow(
		ctx,
		`SELECT a.id, a.slug, a.title, a.description, a.body, a.created_at, a.updated_at, a.user_id, u.username, u.bio, u.image, array_agg(DISTINCT t.title) tags, array_agg(DISTINCT l.user_id) favorited
FROM articles a
JOIN articles_tags at ON a.id = at.article_id
JOIN tags t ON at.tag_id = t.id
LEFT JOIN users u ON a.user_id = u.id
LEFT JOIN likes l ON a.id = l.article_id
WHERE a.slug = $1
GROUP BY a.id, u.id`,
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
		&article.Author.Username,
		&article.Author.Bio,
		&article.Author.Image,
		&article.TagList,
		&article.FavoritedUsersIds,
	)
	if err != nil {
		return entity.ArticleInput{}, err
	}
	return article, nil
}

func (a ArticleRepo) Create(ctx context.Context, article entity.Article) (uint64, error) {
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
		return 0, err
	}
	id, err := a.GetIdBySlug(ctx, article.Slug)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (a ArticleRepo) Update(ctx context.Context, article entity.Article, slug string) error {
	_, err := a.Conn.Exec(
		ctx,
		"UPDATE articles SET slug = $1, title = $2, description = $3, body = $4 WHERE id = $5",
		article.Slug,
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

func (a ArticleRepo) List(ctx context.Context, filter string) ([]entity.ArticleInput, error) {
	sql := `SELECT a.id, a.slug, a.title, a.description, a.body, a.created_at, a.updated_at, a.user_id,
       u.username, u.bio, u.image, 
       array_agg(DISTINCT t.title) tags,
       array_agg(DISTINCT l.user_id) favorited, 
       array_agg(DISTINCT u_fav.username) favoritedNames
FROM articles a
JOIN articles_tags at ON a.id = at.article_id
JOIN tags t ON at.tag_id = t.id
LEFT JOIN users u ON a.user_id = u.id
LEFT JOIN likes l ON a.id = l.article_id
LEFT JOIN users u_fav ON l.user_id = u_fav.id
GROUP BY a.id, u.id`
	sql += filter
	rows, err := a.Conn.Query(
		ctx,
		sql,
	)
	if err != nil {
		return nil, err
	}

	var articles []entity.ArticleInput
	for rows.Next() {
		var article entity.ArticleInput
		err = rows.Scan(
			&article.Id,
			&article.Slug,
			&article.Title,
			&article.Description,
			&article.Body,
			&article.CreatedAt,
			&article.UpdatedAt,
			&article.UserId,
			&article.Author.Username,
			&article.Author.Bio,
			&article.Author.Image,
			&article.TagList,
			&article.FavoritedUsersIds,
			&article.FavoritedUsersUsernames,
		)
		if err != nil && err != pgx.ErrNoRows {
			return nil, err
		}
		articles = append(articles, article)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return articles, nil
}
