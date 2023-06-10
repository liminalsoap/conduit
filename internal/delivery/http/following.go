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
	useCase usecase.Following
	log     logger.Interface
}

func NewFollowingRoutes(handler *gin.RouterGroup, log logger.Interface, uc usecase.Following, mw *middleware.MiddlewareManager) {
	routes := &followingRoutes{uc, log}

	// TODO: update profile output! urgent!
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

	// find user by username & follow
	followingUsername := c.Param("username")
	isFollowing, err := f.useCase.CheckIsFollowing(c.Request.Context(), followingUsername, user.Id)
	if err != nil {
		f.log.Errorf("failed to follow: %s", err)
		errorResponse(c, http.StatusBadRequest, "follow failed")

		return
	}

	profile := user.PrepareProfileOutput(isFollowing)
	c.JSON(http.StatusOK, profile)
}

func (f followingRoutes) Follow(c *gin.Context) {
	userCtx, _ := c.Get("user")
	user := userCtx.(entity.User)

	// find user by username & follow
	followingUsername := c.Param("username")
	err := f.useCase.Follow(c.Request.Context(), followingUsername, user.Id)
	if err != nil {
		f.log.Errorf("failed to follow: %s", err)
		errorResponse(c, http.StatusBadRequest, "follow failed")

		return
	}

	profile := user.PrepareProfileOutput(true)
	c.JSON(http.StatusOK, profile)
}

func (f followingRoutes) Unfollow(c *gin.Context) {
	userCtx, _ := c.Get("user")
	user := userCtx.(entity.User)

	// find user by username & follow
	followingUsername := c.Param("username")
	err := f.useCase.Unfollow(c.Request.Context(), followingUsername, user.Id)
	if err != nil {
		f.log.Errorf("failed to follow: %s", err)
		errorResponse(c, http.StatusBadRequest, "follow failed")

		return
	}

	profile := user.PrepareProfileOutput(false)
	c.JSON(http.StatusOK, profile)
}
