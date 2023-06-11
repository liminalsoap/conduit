package usecase

import "context"

type ArticleTagUseCase struct {
	repo ArticleTagRepo
}

func NewArticleTagUseCase(r ArticleTagRepo) *ArticleTagUseCase {
	return &ArticleTagUseCase{r}
}

func (a ArticleTagUseCase) Add(ctx context.Context, articleId uint64, tagId uint64) error {
	return a.repo.Add(ctx, articleId, tagId)
}

func (a ArticleTagUseCase) AddList(ctx context.Context, articleId uint64, tagIds []uint64) error {
	for _, id := range tagIds {
		err := a.Add(ctx, articleId, id)
		if err != nil {
			return err
		}
	}
	return nil
}
