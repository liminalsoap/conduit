package http

import (
	"conduit-go/internal/usecase"
	"conduit-go/pkg/logger"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, log logger.Interface, uc usecase.UseCases) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	//swagger

	h := handler.Group("/api")
	{
		NewTagRoutes(h, log, uc.Tag)
		NewUserRoutes(h, log, uc.User)
	}
}
