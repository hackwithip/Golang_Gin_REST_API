package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inder231/cms-app/inits"
	"github.com/inder231/cms-app/pkg/models"
	"github.com/inder231/cms-app/pkg/services"
)

func Authenticate (c *gin.Context) {
	var user models.User
	var keycloakUser models.KeycloakUser

	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	
	// Validate token
	// userId, err := utils.VerifyToken(token)

	// Make api call to keycloak to validate token
	resp, err := services.GetKeyclaokUserInfo(token)
	if err!= nil {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token, login again!"})
        return
    }

	email, ok := resp["email"].(string)
	if !ok {
		fmt.Println("ERROR: email field not found or is not a string")
	}

	// Check if user exists in keycloak
	keycloakUserDetails := inits.DB.Where("email = ?", email).Find(&keycloakUser)
	if keycloakUserDetails.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "User with this email does not exists in keycloak table.",
		})
		return
	}
	// Check if user exist in user table
	userDetails := inits.DB.Where("email = ?", email).Find(&user)
	if userDetails.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "User with this email does not exists in user table.",
		})
		return
	}
	// Set in request object
	c.Set("userId", user.ID)
	c.Next()
}