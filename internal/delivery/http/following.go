package http

import (
	"conduit-go/internal/entity"
	"conduit-go/internal/middleware"
	"conduit-go/internal/usecase"
	"conduit-go/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type followingRoutes struct {
	useCase     usecase.Following
	userUseCase usecase.User
	log         logger.Interface
}

func NewFollowingRoutes(handler *gin.RouterGroup, log logger.Interface, uc usecase.Following, userUC usecase.User, mw *middleware.MiddlewareManager) {
	routes := &followingRoutes{uc, userUC, log}

	h := handler.Group("profiles")
	{
		h.GET("/:username", mw.AuthMiddleware, routes.GetProfile)
		h.POST("/:username/follow", mw.AuthMiddleware, routes.Follow)
		h.DELETE("/:username/follow", mw.AuthMiddleware, routes.Unfollow)
	}
}

func (f followingRoutes) GetProfile(c *gin.Context) {
	// TODO: make optional auth
	userCtx, _ := c.Get("user")
	user := userCtx.(entity.User)
	followingUsername := c.Param("username")

	// find user by username & follow
	followingUser, err := f.userUseCase.FindByUsername(c.Request.Context(), followingUsername)
	isFollowing, err := f.useCase.CheckIsFollowing(c.Request.Context(), user.Id, followingUser.Id)
	if err != nil {
		f.log.Errorf("failed to follow: %s", err)
		errorResponse(c, http.StatusBadRequest, "follow failed")

		return
	}

	profile := followingUser.PrepareProfileOutput(isFollowing)
	c.JSON(http.StatusOK, profile)
}

func (f followingRoutes) Follow(c *gin.Context) {
	userCtx, _ := c.Get("user")
	user := userCtx.(entity.User)
	followingUsername := c.Param("username")

	// find user by username & follow
	followingUser, err := f.userUseCase.FindByUsername(c.Request.Context(), followingUsername)
	err = f.useCase.Follow(c.Request.Context(), user.Id, followingUser.Id)
	if err != nil {
		f.log.Errorf("failed to follow: %s", err)
		errorResponse(c, http.StatusBadRequest, "follow failed")

		return
	}

	profile := followingUser.PrepareProfileOutput(true)
	c.JSON(http.StatusOK, profile)
}

func (f followingRoutes) Unfollow(c *gin.Context) {
	userCtx, _ := c.Get("user")
	user := userCtx.(entity.User)
	followingUsername := c.Param("username")

	// find user by username & follow
	followingUser, err := f.userUseCase.FindByUsername(c.Request.Context(), followingUsername)
	err = f.useCase.Unfollow(c.Request.Context(), user.Id, followingUser.Id)

	if err != nil {
		f.log.Errorf("failed to follow: %s", err)
		errorResponse(c, http.StatusBadRequest, "follow failed")

		return
	}

	profile := followingUser.PrepareProfileOutput(false)
	c.JSON(http.StatusOK, profile)
}
