package entity

import (
	"time"
)

type Comment struct {
	Id        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	Body      string
	UserId    uint64
	ArticleId uint64
}

type CommentInput struct {
	Id        uint64    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Body      string    `json:"body"`
	Author    User      `json:"author"`
}

type CommentOutput struct {
	Id        uint64    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Body      string    `json:"body"`
	Author    Profile   `json:"author"`
}

type CommentOutputAlias struct {
	Comment CommentOutput `json:"comment"`
}

type CommentsOutputAlias struct {
	Comment []CommentOutput `json:"comments"`
}

func (c *Comment) PrepareOutput(user User, following bool) CommentOutput {
	return prepareDataForComments(c, user, following)
}

func (c *CommentInput) PrepareOutput() CommentOutput {
	return prepareDataForComments(&Comment{
		Id:        c.Id,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		Body:      c.Body,
	}, User{
		Username: c.Author.Username,
		Bio:      c.Author.Bio,
		Image:    c.Author.Bio,
	}, false)
}

func prepareDataForComments(c *Comment, user User, following bool) CommentOutput {
	return CommentOutput{
		c.Id,
		c.CreatedAt,
		c.UpdatedAt,
		c.Body,
		Profile{
			user.Username,
			user.Bio.String,
			user.Image.String,
			following,
		},
	}
}
