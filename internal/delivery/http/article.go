package http

import (
	"conduit-go/internal/entity"
	"conduit-go/internal/middleware"
	"conduit-go/internal/usecase"
	"conduit-go/pkg/logger"
	"conduit-go/pkg/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
)

type articleRoutes struct {
	useCase          usecase.Article
	followingUseCase usecase.Following
	userUseCase      usecase.User
	likeUseCase      usecase.Like
	commentUseCase   usecase.Comment
	log              logger.Interface
}

func NewArticleRoutes(
	handler *gin.RouterGroup,
	log logger.Interface,
	uc usecase.Article,
	followingUc usecase.Following,
	userUc usecase.User,
	likeUc usecase.Like,
	commentUc usecase.Comment,
	mw *middleware.MiddlewareManager,
) {
	routes := articleRoutes{
		uc,
		followingUc,
		userUc, likeUc,
		commentUc,
		log,
	}

	h := handler.Group("/articles")
	{
		h.POST("/", mw.AuthMiddleware, routes.Create)
		h.GET("/:slug", routes.Get)
		h.PUT("/:slug", mw.AuthMiddleware, routes.Update)
		h.DELETE("/:slug", mw.AuthMiddleware, routes.Delete)

		h.GET("/", routes.List)

		h.POST("/:slug/comments", mw.AuthMiddleware, routes.AddComment)
		h.GET("/:slug/comments", routes.GetComments)
		h.DELETE("/:slug/comments/:id", mw.AuthMiddleware, routes.DeleteComment)

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

	outputArticle := createdArticle.PrepareArticleOutput()
	c.JSON(http.StatusOK, entity.ArticleOutputAlias{Article: outputArticle})
}

func (a articleRoutes) Get(c *gin.Context) {
	slug := c.Param("slug")

	article, err := a.useCase.GetBySlug(c.Request.Context(), slug)
	if err != nil {
		a.log.Errorf("error get article by slug: %s", err)
		errorResponse(c, http.StatusBadRequest, "error get article by slug")

		return
	}

	outputArticle := article.PrepareArticleOutput()
	c.JSON(http.StatusOK, entity.ArticleOutputAlias{Article: outputArticle})
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

	if article.UserId != user.Id {
		err = errors.New("permission denied")
		a.log.Errorf("error get article by slug: %s", err)
		errorResponse(c, http.StatusUnauthorized, err.Error())

		return
	}

	articleEntity := entity.Article{
		Id:          article.Id,
		Slug:        article.Slug,
		Title:       article.Title,
		Description: article.Description,
		Body:        article.Body,
		CreatedAt:   article.CreatedAt,
		UpdatedAt:   article.UpdatedAt,
		UserId:      article.UserId,
	}
	articleEntity.SetInputData(input.Article.Title, input.Article.Description, input.Article.Body)
	if input.Article.Title != "" {
		articleEntity.GenerateSlug()
	}

	if err = a.useCase.Update(c.Request.Context(), articleEntity, slug); err != nil {
		a.log.Errorf("error update: %s", err)
		errorResponse(c, http.StatusBadRequest, "error update")

		return
	}

	article.SetArticleData(articleEntity.Slug, articleEntity.Title, articleEntity.Description, articleEntity.Body)
	article.SetFavorited(user.Id)
	outputArticle := article.PrepareArticleOutput()
	c.JSON(http.StatusOK, entity.ArticleOutputAlias{Article: outputArticle})
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

	article.Favorited = true
	outputArticle := article.PrepareArticleOutput()
	outputArticle.FavoritesCount++
	c.JSON(http.StatusOK, entity.ArticleOutputAlias{Article: outputArticle})
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

	outputArticle := article.PrepareArticleOutput()
	c.JSON(http.StatusOK, entity.ArticleOutputAlias{Article: outputArticle})
}

type ArticlesOutput struct {
	Articles      []entity.ArticleInput `json:"articles"`
	ArticlesCount int                   `json:"articlesCount"`
}

func (a articleRoutes) List(c *gin.Context) {
	tag := c.Query("tag")
	author := c.Query("author")
	favorited := c.Query("favorited")
	limit := c.Query("limit")
	offset := c.Query("offset")

	filter := utils.CreateListFilter(tag, author, favorited, limit, offset)

	articles, err := a.useCase.List(c.Request.Context(), filter)
	if err != nil {
		a.log.Errorf("error get articles: %s", err)
		errorResponse(c, http.StatusBadRequest, "error get articles")

		return
	}
	var outputArticles []entity.ArticleOutput
	for _, article := range articles {
		outputArticles = append(outputArticles, article.PrepareArticleOutput())
	}

	c.JSON(http.StatusOK, entity.ArticlesOutputAlias{Articles: outputArticles, ArticlesCount: len(outputArticles)})
}

type CommentInput struct {
	Comment struct {
		Body string `json:"body"`
	} `json:"comment"`
}

func (a articleRoutes) AddComment(c *gin.Context) {
	userCtx, _ := c.Get("user")
	user := userCtx.(entity.User)

	slug := c.Param("slug")

	var input CommentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		a.log.Errorf("error bind: %s", err)
		errorResponse(c, http.StatusBadRequest, "error bind")

		return
	}

	comment := entity.Comment{
		Body:      input.Comment.Body,
		UserId:    user.Id,
		ArticleId: 0,
	}
	createdComment, err := a.commentUseCase.Add(c.Request.Context(), slug, comment)
	if err != nil {
		a.log.Errorf("error add comment: %s", err)
		errorResponse(c, http.StatusBadRequest, "error add comment")

		return
	}

	commentAuthor, err := a.userUseCase.GetUser(c.Request.Context(), user.Id)
	if err != nil {
		a.log.Errorf("error find author: %s", err)
		errorResponse(c, http.StatusBadRequest, "error find author")

		return
	}

	output := createdComment.PrepareOutput(commentAuthor, false)
	c.JSON(http.StatusOK, entity.CommentOutputAlias{Comment: output})
}

func (a articleRoutes) GetComments(c *gin.Context) {
	slug := c.Param("slug")

	comments, err := a.commentUseCase.GetByArticleId(c.Request.Context(), slug)
	if err != nil {
		a.log.Errorf("error get comments: %s", err)
		errorResponse(c, http.StatusBadRequest, "error get comments")

		return
	}

	var outputComments []entity.CommentOutput
	for _, comment := range comments {
		outputComments = append(outputComments, comment.PrepareOutput())
	}

	// TODO: fix following output
	c.JSON(http.StatusOK, entity.CommentsOutputAlias{Comment: outputComments})
}

func (a articleRoutes) DeleteComment(c *gin.Context) {
	userCtx, _ := c.Get("user")
	user := userCtx.(entity.User)

	id := c.Param("id")

	comment, err := a.commentUseCase.GetById(c.Request.Context(), cast.ToUint64(id))
	if err != nil {
		a.log.Errorf("comment doesn't exist: %s", err)
		errorResponse(c, http.StatusBadRequest, "comment doesn't exist")

		return
	}

	if user.Id != comment.UserId {
		a.log.Errorf("comment doesn't exist: %s", errors.New("user isn't owner of comment"))
		errorResponse(c, http.StatusUnauthorized, "user isn't owner of comment")

		return
	}

	if err = a.commentUseCase.Delete(c.Request.Context(), cast.ToUint64(id)); err != nil {
		a.log.Errorf("error delete comment: %s", err)
		errorResponse(c, http.StatusBadRequest, "error delete comment")

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
