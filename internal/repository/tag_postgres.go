package repository

import (
	"conduit-go/internal/entity"
	"conduit-go/pkg/postgres"
	"context"
)

type TagRepo struct {
	*postgres.Postgres
}

func NewTagRepo(pg *postgres.Postgres) *TagRepo {
	return &TagRepo{pg}
}

func (t TagRepo) GetTags(ctx context.Context) (*[]entity.Tag, error) {
	rows, err := t.Conn.Query(ctx, "SELECT * FROM tags")
	if err != nil {
		return nil, err
	}

	var tags []entity.Tag
	for rows.Next() {
		var tag entity.Tag
		err := rows.Scan(
			&tag.Id,
			&tag.Title,
		)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	if err := rows.Err(); err != nil {

	}

	return &tags, err
}

func (t TagRepo) GetByTitle(ctx context.Context, title string) (uint64, error) {
	var id uint64
	err := t.Conn.QueryRow(ctx, "SELECT id FROM tags WHERE title = $1", title).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (t TagRepo) GetByTitles(ctx context.Context, titles []string) ([]uint64, error) {
	rows, err := t.Conn.Query(ctx, "SELECT id FROM tags WHERE title = ANY($1)", titles)
	if err != nil {
		return []uint64{}, err
	}

	var ids []uint64
	for rows.Next() {
		var id uint64
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return ids, nil
}

func (t TagRepo) GetByIds(ctx context.Context, ids []uint64) ([]string, error) {
	rows, err := t.Conn.Query(ctx, "SELECT title FROM tags WHERE id = ANY($1)", ids)
	if err != nil {
		return []string{}, err
	}

	var titles []string
	for rows.Next() {
		var title string
		err = rows.Scan(&title)
		if err != nil {
			return nil, err
		}
		titles = append(titles, title)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return titles, nil
}
