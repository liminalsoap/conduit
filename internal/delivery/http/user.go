package http

import (
	"conduit-go/config"
	"conduit-go/internal/entity"
	"conduit-go/internal/middleware"
	"conduit-go/internal/usecase"
	"conduit-go/pkg/logger"
	"conduit-go/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userRoutes struct {
	useCase usecase.User
	log     logger.Interface
	cfg     *config.Config
}

const tokenClaim = "email"

func NewUserRoutes(handler *gin.RouterGroup, log logger.Interface, uc usecase.User, cfg *config.Config, mw *middleware.MiddlewareManager) {
	routes := userRoutes{uc, log, cfg}

	h := handler.Group("/users")
	{
		h.POST("/", routes.Register)
		h.POST("/login", routes.Login)
	}
	handler.GET("/user", mw.AuthMiddleware, routes.GetCurrentUser)
}

type createInput struct {
	User struct {
		Email    string `json:"email" binding:"required"`
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	} `json:"user"`
}

func (u userRoutes) Register(c *gin.Context) {
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
	if err := user.HashPassword(); err != nil {
		u.log.Errorf("error hash: %s", err)
		errorResponse(c, http.StatusBadRequest, "error hash")

		return
	}

	usr, err := u.useCase.Create(c.Request.Context(), user)
	if err != nil {
		u.log.Errorf("error create: %s", err)
		errorResponse(c, http.StatusBadRequest, "error create")

		return
	}

	output := usr.PrepareOutput()
	c.JSON(http.StatusOK, output)
}

type loginInput struct {
	User struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	} `json:"user"`
}

func (u userRoutes) Login(c *gin.Context) {
	// bind input
	var input loginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		u.log.Errorf("error bind: %s", err)
		errorResponse(c, http.StatusBadRequest, "error bind")

		return
	}

	// find by email
	user, err := u.useCase.FindByEmail(c.Request.Context(), input.User.Email)
	if err != nil {
		u.log.Errorf("error find user: %s", err)
		errorResponse(c, http.StatusBadRequest, "error find")

		return
	}

	// compare passwords
	if err := user.ComparePassword(input.User.Password); err != nil {
		u.log.Errorf("password is invalid: %s", err)
		errorResponse(c, http.StatusUnauthorized, "password is invalid")

		return
	}

	// generate jwt
	token, err := utils.NewToken(u.cfg.Http.Secret, user.Email)
	if err != nil {
		u.log.Errorf("failed to generate token: %s", err)
		errorResponse(c, http.StatusUnauthorized, "failed to generate token")

		return
	}

	// respond
	output := user.PrepareOutput()
	output.Token = token
	c.JSON(http.StatusOK, output)
}

func (u userRoutes) GetCurrentUser(c *gin.Context) {
	output, _ := c.Get("user")
	c.JSON(http.StatusOK, output)
}
