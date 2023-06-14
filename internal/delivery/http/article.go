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
	userUseCase      usecase.User
	likeUseCase      usecase.Like
	log              logger.Interface
}

func NewArticleRoutes(
	handler *gin.RouterGroup,
	log logger.Interface,
	uc usecase.Article,
	followingUc usecase.Following,
	userUc usecase.User,
	likeUc usecase.Like,
	mw *middleware.MiddlewareManager,
) {
	routes := articleRoutes{uc, followingUc, userUc, likeUc, log}

	h := handler.Group("/articles")
	{
		h.POST("/", mw.AuthMiddleware, routes.Create)
		h.GET("/:slug", routes.Get)
		h.PUT("/:slug", mw.AuthMiddleware, routes.Update)
		h.DELETE("/:slug", mw.AuthMiddleware, routes.Delete)

		h.POST("/:slug/favorite", mw.AuthMiddleware, routes.Favorite)
		h.DELETE("/:slug/favorite", mw.AuthMiddleware, routes.Unfavorite)
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

func (a articleRoutes) Get(c *gin.Context) {
	slug := c.Param("slug")

	article, err := a.useCase.GetBySlug(c.Request.Context(), slug)
	if err != nil {
		a.log.Errorf("error get article by slug: %s", err)
		errorResponse(c, http.StatusBadRequest, "error get article by slug")

		return
	}

	tagList, err := a.useCase.GetTagList(c.Request.Context(), article.Id)
	if err != nil {
		a.log.Errorf("error get tags: %s", err)
		errorResponse(c, http.StatusBadRequest, "error get tags")

		return
	}

	author, err := a.userUseCase.GetUser(c.Request.Context(), article.UserId)
	if err != nil {
		a.log.Errorf("failed found user: %s", err)
		errorResponse(c, http.StatusBadRequest, "failed found user")

		return
	}

	count, err := a.likeUseCase.Count(c.Request.Context(), article.Id, author.Id)
	if err != nil {
		a.log.Errorf("error count likes: %s", err)
		errorResponse(c, http.StatusBadRequest, "error count likes")

		return
	}

	outputArticle := article.PrepareOutput(tagList, false, count, author.PrepareReuseProfileOutput(false))
	c.JSON(http.StatusOK, outputArticle)
}

type updateArticleInput struct {
	Article struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Body        string `json:"body"`
	} `json:"article"`
}

func (a articleRoutes) Update(c *gin.Context) {
	userCtx, _ := c.Get("user")
	user := userCtx.(entity.User)

	slug := c.Param("slug")

	var input updateArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		a.log.Errorf("error bind: %s", err)
		errorResponse(c, http.StatusBadRequest, "error bind")

		return
	}

	article, err := a.useCase.GetBySlug(c.Request.Context(), slug)
	if err != nil {
		a.log.Errorf("error get article by slug: %s", err)
		errorResponse(c, http.StatusBadRequest, "error get article by slug")

		return
	}

	article.SetInputData(input.Article.Title, input.Article.Description, input.Article.Body)

	if err = a.useCase.Update(c.Request.Context(), article, slug); err != nil {
		a.log.Errorf("error update: %s", err)
		errorResponse(c, http.StatusBadRequest, "error update")

		return
	}
	if input.Article.Title != "" {
		article.GenerateSlug()
	}

	tagList, err := a.useCase.GetTagList(c.Request.Context(), article.Id)
	if err != nil {
		a.log.Errorf("error get tags: %s", err)
		errorResponse(c, http.StatusBadRequest, "error get tags")

		return
	}

	count, err := a.likeUseCase.Count(c.Request.Context(), article.Id, user.Id)
	if err != nil {
		a.log.Errorf("error count likes: %s", err)
		errorResponse(c, http.StatusBadRequest, "error count likes")

		return
	}

	outputArticle := article.PrepareOutput(tagList, false, count, user.PrepareReuseProfileOutput(DefaultMyselfFavorited))
	c.JSON(http.StatusOK, outputArticle)
}

func (a articleRoutes) Delete(c *gin.Context) {
	slug := c.Param("slug")

	err := a.useCase.DeleteBySlug(c.Request.Context(), slug)
	if err != nil {
		a.log.Errorf("error delete: %s", err)
		errorResponse(c, http.StatusBadRequest, "error delete")

		return
	}

	c.Status(http.StatusOK)
}

func (a articleRoutes) Favorite(c *gin.Context) {
	userCtx, _ := c.Get("user")
	user := userCtx.(entity.User)

	slug := c.Param("slug")

	article, err := a.useCase.GetBySlug(c.Request.Context(), slug)
	if err != nil {
		a.log.Errorf("error get article: %s", err)
		errorResponse(c, http.StatusBadRequest, "error get article")

		return
	}

	if err = a.likeUseCase.Favorite(c.Request.Context(), article.Id, user.Id); err != nil {
		a.log.Errorf("error like article: %s", err)
		errorResponse(c, http.StatusBadRequest, "error like article")

		return
	}

	tagList, err := a.useCase.GetTagList(c.Request.Context(), article.Id)
	if err != nil {
		a.log.Errorf("error get tags: %s", err)
		errorResponse(c, http.StatusBadRequest, "error get tags")

		return
	}

	count, err := a.likeUseCase.Count(c.Request.Context(), article.Id, user.Id)
	if err != nil {
		a.log.Errorf("error count likes: %s", err)
		errorResponse(c, http.StatusBadRequest, "error count likes")

		return
	}

	outputArticle := article.PrepareOutput(tagList, true, count, user.PrepareReuseProfileOutput(DefaultMyselfFavorited))
	c.JSON(http.StatusOK, outputArticle)
}

func (a articleRoutes) Unfavorite(c *gin.Context) {
	userCtx, _ := c.Get("user")
	user := userCtx.(entity.User)

	slug := c.Param("slug")

	article, err := a.useCase.GetBySlug(c.Request.Context(), slug)
	if err != nil {
		a.log.Errorf("error get article: %s", err)
		errorResponse(c, http.StatusBadRequest, "error get article")

		return
	}

	if err = a.likeUseCase.Unfavorite(c.Request.Context(), article.Id, user.Id); err != nil {
		a.log.Errorf("error like article: %s", err)
		errorResponse(c, http.StatusBadRequest, "error like article")

		return
	}

	tagList, err := a.useCase.GetTagList(c.Request.Context(), article.Id)
	if err != nil {
		a.log.Errorf("error get tags: %s", err)
		errorResponse(c, http.StatusBadRequest, "error get tags")

		return
	}

	count, err := a.likeUseCase.Count(c.Request.Context(), article.Id, user.Id)
	if err != nil {
		a.log.Errorf("error count likes: %s", err)
		errorResponse(c, http.StatusBadRequest, "error count likes")

		return
	}

	outputArticle := article.PrepareOutput(tagList, false, count, user.PrepareReuseProfileOutput(DefaultMyselfFavorited))
	c.JSON(http.StatusOK, outputArticle)
}
