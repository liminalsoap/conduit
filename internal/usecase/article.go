package usecase

import (
	"conduit-go/internal/entity"
	"context"
)

type ArticleUseCase struct {
	repo           ArticleRepo
	tagRepo        TagUseCase
	articleTagRepo ArticleTagUseCase
}

func NewArticleUseCase(r ArticleRepo, tagU TagUseCase, articleTagU ArticleTagUseCase) *ArticleUseCase {
	return &ArticleUseCase{r, tagU, articleTagU}
}

func (a ArticleUseCase) GetBySlug(ctx context.Context, slug string) (entity.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleUseCase) Create(ctx context.Context, article entity.Article, titles []string) (entity.Article, error) {
	titlesIds, err := a.tagRepo.GetByTitles(ctx, titles)
	if err != nil {
		return entity.Article{}, err
	}
	err = a.articleTagRepo.AddList(ctx, article.Id, titlesIds)
	if err != nil {
		return entity.Article{}, err
	}
	return a.repo.Create(ctx, article)
}

func (a ArticleUseCase) Update(ctx context.Context, slug string) (entity.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleUseCase) DeleteBySlug(ctx context.Context, slug string) (entity.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleUseCase) List(ctx context.Context, slug string) ([]entity.Article, error) {
	//TODO implement me
	panic("implement me")
}
