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
