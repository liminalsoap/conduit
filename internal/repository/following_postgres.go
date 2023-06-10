package repository

import (
	"conduit-go/pkg/postgres"
	"context"
	"errors"
)

type FollowingRepo struct {
	*postgres.Postgres
}

func NewFollowingRepo(pg *postgres.Postgres) *FollowingRepo {
	return &FollowingRepo{pg}
}

func (r FollowingRepo) HelperIsValid(ctx context.Context, followingUserId, followedUserId uint64, isFollow bool) error {
	if followingUserId == followedUserId {
		return errors.New("ids must be different")
	}
	isFollowing, err := r.CheckIsFollowing(ctx, followingUserId, followedUserId)
	if err != nil {
		return err
	}
	if isFollowing && isFollow || !isFollowing && !isFollow {
		return errors.New("user already followed")
	}
	return nil
}

func (r FollowingRepo) Follow(ctx context.Context, followingUserId, followedUserId uint64) error {
	if err := r.HelperIsValid(ctx, followingUserId, followedUserId, true); err != nil {
		return err
	}
	_, err := r.Conn.Exec(
		ctx,
		"INSERT INTO follows(following_user_id, followed_user_id) VALUES ($1, $2)",
		followingUserId,
		followedUserId,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r FollowingRepo) Unfollow(ctx context.Context, followingUserId, followedUserId uint64) error {
	if err := r.HelperIsValid(ctx, followingUserId, followedUserId, false); err != nil {
		return err
	}
	_, err := r.Conn.Exec(
		ctx,
		"DELETE FROM follows WHERE following_user_id = $1 and followed_user_id = $2",
		followingUserId,
		followedUserId,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r FollowingRepo) CheckIsFollowing(ctx context.Context, followingUserId, followedUserId uint64) (bool, error) {
	var id uint64
	rows, err := r.Conn.Query(
		ctx,
		"SELECT id FROM follows WHERE following_user_id = $1 and followed_user_id = $2",
		followingUserId,
		followedUserId,
	)
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return false, err
		}
	}
	if rows.Err() != nil {
		return false, err
	}
	if id == 0 {
		return false, nil
	}
	return true, nil
}
