package http

import (
	"conduit-go/config"
	"conduit-go/internal/usecase"
	"conduit-go/pkg/logger"
	"github.com/gin-gonic/gin"
)

func NewRouter(cfg *config.Config, handler *gin.Engine, log logger.Interface, uc usecase.UseCases) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	//swagger

	h := handler.Group("/api")
	{
		NewTagRoutes(h, log, uc.Tag)
		NewUserRoutes(cfg, h, log, uc.User)
	}
}
