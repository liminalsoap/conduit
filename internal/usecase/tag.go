package usecase

import (
	"conduit-go/internal/entity"
	"context"
)

type TagUseCase struct {
	repo TagRepo
}

func NewTagUseCase(r TagRepo) *TagUseCase {
	return &TagUseCase{r}
}

func (t TagUseCase) List(ctx context.Context) (*[]entity.Tag, error) {
	return t.repo.GetTags(ctx)
}
