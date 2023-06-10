package usecase

import (
	"conduit-go/internal/entity"
	"context"
)

type UserUseCase struct {
	repo UserRepo
}

func NewUserUseCase(r UserRepo) *UserUseCase {
	return &UserUseCase{r}
}

func (uc UserUseCase) Create(ctx context.Context, user entity.User) (entity.User, error) {
	return uc.repo.Create(ctx, user)
}

func (uc UserUseCase) GetUser(ctx context.Context, id uint64) (entity.User, error) {
	return uc.repo.GetById(ctx, id)
}

func (uc UserUseCase) Update(ctx context.Context, user entity.User) (entity.User, error) {
	return uc.repo.Update(ctx, user)
}

func (uc UserUseCase) FindByEmail(ctx context.Context, email string) (entity.User, error) {
	return uc.repo.FindByEmail(ctx, email)
}

func (uc UserUseCase) FindByUsername(ctx context.Context, username string) (entity.User, error) {
	return uc.repo.FindByEmail(ctx, username)
}
