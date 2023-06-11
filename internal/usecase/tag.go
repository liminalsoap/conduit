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

func (uc TagUseCase) List(ctx context.Context) (*[]entity.Tag, error) {
	return uc.repo.GetTags(ctx)
}

func (uc TagUseCase) GetByTitle(ctx context.Context, title string) (uint64, error) {
	return uc.repo.GetByTitle(ctx, title)
}

func (uc TagUseCase) GetByTitles(ctx context.Context, titles []string) ([]uint64, error) {
	return uc.repo.GetByTitles(ctx, titles)
}
