package entity

import "time"

type Following struct {
	Id              uint64
	FollowingUserId uint64
	FollowedUserId  uint64
	CreatedAt       time.Time
}
