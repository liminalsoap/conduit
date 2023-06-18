package usecase

import (
	"conduit-go/internal/entity"
	"context"
	"fmt"
)

type CommentUseCase struct {
	repo           CommentRepo
	articleUseCase *ArticleUseCase
}

func NewCommentUseCase(r CommentRepo, articleUC *ArticleUseCase) *CommentUseCase {
	return &CommentUseCase{r, articleUC}
}

func (c CommentUseCase) Add(ctx context.Context, slug string, comment entity.Comment) (entity.Comment, error) {
	article, err := c.articleUseCase.GetBySlug(ctx, slug)
	if err != nil {
		return entity.Comment{}, err
	}
	return c.repo.Add(ctx, article.Id, comment)
}

func (c CommentUseCase) GetByArticleId(ctx context.Context, slug string) ([]entity.CommentInput, error) {
	article, err := c.articleUseCase.GetBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}
	return c.repo.GetByArticleId(ctx, article.Id)
}

func (c CommentUseCase) Delete(ctx context.Context, id uint64) error {
	return c.repo.Delete(ctx, id)
}

func (c CommentUseCase) GetById(ctx context.Context, id uint64) (entity.Comment, error) {
	fmt.Println(id)
	return c.repo.GetById(ctx, id)
}
