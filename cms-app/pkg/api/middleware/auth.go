package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inder231/cms-app/pkg/utils"
)

func Authenticate (c *gin.Context) {

	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	
	// Validate token
	userId, err := utils.VerifyToken(token)
	
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token, login again!"})
		return
	}
	// Set in request object
	c.Set("userId", userId)
	
	c.Next()
}