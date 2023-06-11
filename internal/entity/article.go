package entity

import "time"

type Article struct {
	Id          uint64
	Slug        string
	Title       string
	Description string
	Body        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserId      uint64
}
