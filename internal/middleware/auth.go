package middleware

import (
	"conduit-go/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

const authHeader = "Authorization"
const tokenClaim = "email"

func (mw *MiddlewareManager) AuthMiddleware(c *gin.Context) {
	token := c.GetHeader(authHeader)

	email, err := utils.GetClaimByTokenAndName(mw.cfg.Http.Secret, token, tokenClaim)
	if err != nil {
		mw.log.Errorf("failed to parse token: %s", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "error auth"})

		return
	}

	user, err := mw.userUC.FindByEmail(c.Request.Context(), email)
	if err != nil {
		mw.log.Errorf("error find user: %s", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "error auth"})

		return
	}

	// respond
	output := user.PrepareOutput()
	output.Token = token
	c.Set("user", output)
	c.Next()
}
