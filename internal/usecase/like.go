package usecase

import "context"

type LikeUseCase struct {
	repo LikeRepo
}

func NewLikeUseCase(r LikeRepo) *LikeUseCase {
	return &LikeUseCase{r}
}

func (l LikeUseCase) Favorite(ctx context.Context, articleId, userId uint64) error {
	return l.repo.Add(ctx, articleId, userId)
}

func (l LikeUseCase) Unfavorite(ctx context.Context, articleId, userId uint64) error {
	return l.repo.Delete(ctx, articleId, userId)
}

func (l LikeUseCase) Count(ctx context.Context, articleId uint64) (uint64, error) {
	return l.repo.Count(ctx, articleId)
}
