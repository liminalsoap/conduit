package usecase

import (
	"conduit-go/internal/entity"
	"context"
	"errors"
)

type ArticleUseCase struct {
	repo           ArticleRepo
	tagRepo        *TagUseCase
	articleTagRepo *ArticleTagUseCase
}

func NewArticleUseCase(r ArticleRepo, tagU *TagUseCase, articleTagU *ArticleTagUseCase) *ArticleUseCase {
	return &ArticleUseCase{r, tagU, articleTagU}
}

func (a ArticleUseCase) GetBySlug(ctx context.Context, slug string) (entity.Article, error) {
	return a.repo.GetBySlug(ctx, slug)
}

func (a ArticleUseCase) Create(ctx context.Context, article entity.Article, titles []string) (entity.Article, error) {
	createdArticle, err := a.repo.Create(ctx, article)
	if err != nil {
		return entity.Article{}, err
	}
	if titles != nil {
		titlesIds, err := a.tagRepo.GetByTitles(ctx, titles)
		if err != nil {
			return entity.Article{}, err
		}
		if titlesIds == nil {
			return entity.Article{}, errors.New("tags incorrect")
		}
		err = a.articleTagRepo.AddList(ctx, createdArticle.Id, titlesIds)
		if err != nil {
			return entity.Article{}, err
		}
	}
	return createdArticle, nil
}

func (a ArticleUseCase) Update(ctx context.Context, article entity.Article, slug string) error {
	return a.repo.Update(ctx, article, slug)
}

func (a ArticleUseCase) DeleteBySlug(ctx context.Context, slug string) error {
	return a.repo.DeleteBySlug(ctx, slug)
}

func (a ArticleUseCase) List(ctx context.Context, slug string) ([]entity.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleUseCase) GetTagList(ctx context.Context, id uint64) ([]string, error) {
	tagIds, err := a.articleTagRepo.GetTagsByArticleId(ctx, id)
	if err != nil {
		return nil, err
	}
	tags, err := a.tagRepo.GetByIds(ctx, tagIds)
	if err != nil {
		return nil, err
	}
	return tags, nil
}
