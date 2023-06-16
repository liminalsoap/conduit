package entity

import (
	"database/sql"
	"github.com/spf13/cast"
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

type ArticleInput struct {
	Id                 uint64          `json:"id,omitempty"`
	Slug               string          `json:"slug"`
	Title              string          `json:"title"`
	Description        string          `json:"description"`
	Body               string          `json:"body"`
	TagList            []string        `json:"tagList"`
	CreatedAt          time.Time       `json:"createdAt"`
	UpdatedAt          time.Time       `json:"updatedAt"`
	UserId             uint64          `json:"userId,omitempty"`
	Favorited          bool            `json:"favorited"`
	FavoritedUsersList []sql.NullInt64 `json:"favoritedUsersList,omitempty"`
	FavoritesCount     uint64          `json:"favoritesCount"`
	Author             User            `json:"author"`
}

type ArticleOutput struct {
	Slug           string    `json:"slug"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Body           string    `json:"body"`
	TagList        []string  `json:"tagList"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	UserId         uint64    `json:"userId,omitempty"`
	Favorited      bool      `json:"favorited"`
	FavoritesCount int       `json:"favoritesCount"`
	Author         Profile   `json:"author"`
}

type ArticleOutputAlias struct {
	Article ArticleOutput `json:"article"`
}
type ArticlesOutputAlias struct {
	Articles      []ArticleOutput `json:"articles"`
	ArticlesCount int             `json:"articlesCount"`
}

func (a *ArticleInput) SetFavorited(userId uint64) {
	for _, favoritedId := range a.FavoritedUsersList {
		if favoritedId.Int64 == cast.ToInt64(userId) {
			a.Favorited = true
		}
	}
}

func (a *ArticleInput) PrepareArticleOutput() ArticleOutput {
	favorites := 0
	if a.FavoritedUsersList[0].Valid {
		favorites = len(a.FavoritedUsersList)
	}
	return ArticleOutput{
		a.Slug,
		a.Title,
		a.Description,
		a.Body,
		a.TagList,
		a.CreatedAt,
		a.UpdatedAt,
		a.UserId,
		a.Favorited,
		favorites,
		Profile{
			Username:  a.Author.Username,
			Bio:       a.Author.Bio.String,
			Image:     a.Author.Bio.String,
			Following: false,
		},
	}
}

func (a *ArticleInput) SetArticleData(slug, title, description, body string) {
	a.Slug = slug
	a.Title = title
	a.Description = description
	a.Body = body
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
