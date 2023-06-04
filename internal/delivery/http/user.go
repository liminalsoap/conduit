package http

import (
	"conduit-go/internal/entity"
	"conduit-go/internal/usecase"
	"conduit-go/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userRoutes struct {
	useCase usecase.User
	log     logger.Interface
}

func NewUserRoutes(handler *gin.RouterGroup, log logger.Interface, uc usecase.User) {
	routes := userRoutes{uc, log}

	h := handler.Group("/users")
	{
		h.POST("/", routes.Create)
	}
}

type createInput struct {
	User struct {
		Email    string `json:"email" binding:"required"`
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	} `json:"user"`
}

func (u userRoutes) Create(c *gin.Context) {
	var input createInput
	if err := c.ShouldBindJSON(&input); err != nil {
		u.log.Errorf("error bind: %s", err)
		errorResponse(c, http.StatusBadRequest, "error bind")
		return
	}

	user := entity.User{
		Email:    input.User.Email,
		Username: input.User.Username,
		Password: input.User.Password,
	}
	usr, err := u.useCase.Create(c.Request.Context(), user)
	if err != nil {
		u.log.Errorf("error create: %s", err)
		errorResponse(c, http.StatusBadRequest, "error create")
		return
	}

	c.JSON(http.StatusOK, usr)
}
