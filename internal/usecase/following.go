package usecase

import (
	"context"
)

type FollowingUseCase struct {
	repo     FollowingRepo
	userRepo UserRepo
}

func NewFollowUseCase(r FollowingRepo, userR UserRepo) *FollowingUseCase {
	return &FollowingUseCase{r, userR}
}

func (uc FollowingUseCase) Follow(ctx context.Context, followingUsername string, followedUserId uint64) error {
	user, err := uc.userRepo.FindByUsername(ctx, followingUsername)
	if err != nil {
		return err
	}
	return uc.repo.Follow(ctx, user.Id, followedUserId)
}

func (uc FollowingUseCase) Unfollow(ctx context.Context, followingUsername string, followedUserId uint64) error {
	user, err := uc.userRepo.FindByUsername(ctx, followingUsername)
	if err != nil {
		return err
	}
	return uc.repo.Unfollow(ctx, user.Id, followedUserId)
}

func (uc FollowingUseCase) CheckIsFollowing(ctx context.Context, followingUsername string, followedUserId uint64) (bool, error) {
	user, err := uc.userRepo.FindByUsername(ctx, followingUsername)
	if err != nil {
		return false, err
	}
	return uc.repo.CheckIsFollowing(ctx, user.Id, followedUserId)
}
