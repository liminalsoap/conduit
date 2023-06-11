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

func (uc FollowingUseCase) Follow(ctx context.Context, followingUserId uint64, followedUserId uint64) error {
	return uc.repo.Follow(ctx, followingUserId, followedUserId)
}

func (uc FollowingUseCase) Unfollow(ctx context.Context, followingUserId uint64, followedUserId uint64) error {
	return uc.repo.Unfollow(ctx, followingUserId, followedUserId)
}

func (uc FollowingUseCase) CheckIsFollowing(ctx context.Context, followingUserId uint64, followedUserId uint64) (bool, error) {
	return uc.repo.CheckIsFollowing(ctx, followingUserId, followedUserId)
}
