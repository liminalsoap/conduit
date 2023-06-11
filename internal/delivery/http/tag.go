package http

import (
	"conduit-go/internal/usecase"
	"conduit-go/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type tagRoutes struct {
	useCase usecase.Tag
	log     logger.Interface
}

func NewTagRoutes(handler *gin.RouterGroup, log logger.Interface, uc usecase.Tag) {
	routes := &tagRoutes{uc, log}

	handler.POST("/tags", routes.List)
}

type TagsOutput struct {
	Tags []string `json:"tags"`
}

func (r tagRoutes) List(c *gin.Context) {
	tags, err := r.useCase.List(c.Request.Context())
	if err != nil {
		r.log.Errorf("route err: %s", err)
		errorResponse(c, http.StatusInternalServerError, "internal error")

		return
	}
	var tagsOutput TagsOutput
	for _, tag := range *tags {
		tagsOutput.Tags = append(tagsOutput.Tags, tag.Title)
	}
	c.JSON(http.StatusOK, tagsOutput)
}
