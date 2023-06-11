package http

import (
	"conduit-go/config"
	"conduit-go/internal/middleware"
	"conduit-go/internal/usecase"
	"conduit-go/pkg/logger"
	"github.com/gin-gonic/gin"
)

const authHeader = "Authorization"

func NewRouter(handler *gin.Engine, log logger.Interface, uc usecase.UseCases, cfg *config.Config, mw *middleware.MiddlewareManager) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	//swagger

	h := handler.Group("/api")
	{
		NewTagRoutes(h, log, uc.Tag)
		NewUserRoutes(h, log, uc.User, cfg, mw)
		NewFollowingRoutes(h, log, uc.Following, uc.User, mw)
	}
}
