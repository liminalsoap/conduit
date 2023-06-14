package entity

import (
	"strings"
	"time"
)

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

type ArticleOutput struct {
	Slug           string    `json:"slug"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Body           string    `json:"body"`
	TagList        []string  `json:"tagList"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Favorited      bool      `json:"favorited"`
	FavoritesCount int       `json:"favoritesCount"`
	Author         Profile   `json:"author"`
}

type ArticleOutputName struct {
	ArticleOutput `json:"article"`
}

func (a *Article) PrepareOutput(tagList []string, favorited bool, favoritesCount int, author Profile) ArticleOutputName {
	article := ArticleOutput{
		a.Slug,
		a.Title,
		a.Description,
		a.Body,
		tagList,
		a.CreatedAt,
		a.UpdatedAt,
		favorited,
		favoritesCount,
		author,
	}
	return ArticleOutputName{article}
}

func (a *Article) GenerateSlug() {
	lowerTitle := strings.ToLower(a.Title)
	words := strings.Split(lowerTitle, " ")
	slug := strings.Join(words, "-")
	a.Slug = slug
}

func (a *Article) SetInputData(title, description, body string) {
	if title != "" {
		a.Title = title
	}
	if description != "" {
		a.Description = description
	}
	if body != "" {
		a.Body = body
	}
}
