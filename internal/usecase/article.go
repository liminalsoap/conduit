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

func (a ArticleUseCase) GetBySlug(ctx context.Context, slug string) (entity.ArticleInput, error) {
	return a.repo.GetBySlug(ctx, slug)
}

func (a ArticleUseCase) Create(ctx context.Context, article entity.Article, titles []string) (entity.ArticleInput, error) {
	articleId, err := a.repo.Create(ctx, article)
	if err != nil {
		return entity.ArticleInput{}, err
	}
	if titles != nil {
		titlesIds, err := a.tagRepo.GetByTitles(ctx, titles)
		if err != nil {
			return entity.ArticleInput{}, err
		}
		if titlesIds == nil {
			return entity.ArticleInput{}, errors.New("tags incorrect")
		}
		err = a.articleTagRepo.AddList(ctx, articleId, titlesIds)
		if err != nil {
			return entity.ArticleInput{}, err
		}
	}
	createdArticle, err := a.GetBySlug(ctx, article.Slug)
	if err != nil {
		return entity.ArticleInput{}, err
	}
	return createdArticle, nil
}

func (a ArticleUseCase) Update(ctx context.Context, article entity.Article, slug string) error {
	return a.repo.Update(ctx, article, slug)
}

func (a ArticleUseCase) DeleteBySlug(ctx context.Context, slug string) error {
	return a.repo.DeleteBySlug(ctx, slug)
}

func (a ArticleUseCase) List(ctx context.Context, filter string) ([]entity.ArticleInput, error) {
	return a.repo.List(ctx, filter)
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
