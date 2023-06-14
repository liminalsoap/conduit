package http

import (
	"conduit-go/internal/entity"
	"conduit-go/internal/middleware"
	"conduit-go/internal/usecase"
	"conduit-go/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DefaultMyselfFavorited
// can a user be subscribed to himself
// true - yes, false - no
const DefaultMyselfFavorited = false

type articleRoutes struct {
	useCase          usecase.Article
	followingUseCase usecase.Following
	log              logger.Interface
}

func NewArticleRoutes(handler *gin.RouterGroup, log logger.Interface, uc usecase.Article, fUc usecase.Following, mw *middleware.MiddlewareManager) {
	routes := articleRoutes{uc, fUc, log}

	h := handler.Group("/articles")
	{
		h.POST("/", mw.AuthMiddleware, routes.Create)
	}
}

type createArticleInput struct {
	Article struct {
		Title       string   `json:"title" binding:"required"`
		Description string   `json:"description" binding:"required"`
		Body        string   `json:"body" binding:"required"`
		TagList     []string `json:"tagList"`
	} `json:"article"`
}

func (a articleRoutes) Create(c *gin.Context) {
	userCtx, _ := c.Get("user")
	user := userCtx.(entity.User)

	var input createArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		a.log.Errorf("error bind: %s", err)
		errorResponse(c, http.StatusBadRequest, "error bind")

		return
	}

	article := entity.Article{Title: input.Article.Title, Description: input.Article.Description, Body: input.Article.Body}
	article.GenerateSlug()
	article.UserId = user.Id

	createdArticle, err := a.useCase.Create(c.Request.Context(), article, input.Article.TagList)
	if err != nil {
		a.log.Errorf("error create: %s", err)
		errorResponse(c, http.StatusBadRequest, "error create")

		return
	}

	tagList, err := a.useCase.GetTagList(c.Request.Context(), createdArticle.Id)
	if err != nil {
		a.log.Errorf("error get tags: %s", err)
		errorResponse(c, http.StatusBadRequest, "error get tags")

		return
	}

	outputArticle := createdArticle.PrepareOutput(tagList, false, 0, user.PrepareReuseProfileOutput(DefaultMyselfFavorited))
	c.JSON(http.StatusOK, outputArticle)
}
