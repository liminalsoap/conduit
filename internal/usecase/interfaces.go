package usecase

import (
	"conduit-go/internal/entity"
	"context"
)

type UseCases struct {
	Tag
}

type Tag interface {
	List(ctx context.Context) (*[]entity.Tag, error)
}

type TagRepo interface {
	GetTags(ctx context.Context) (*[]entity.Tag, error)
}
