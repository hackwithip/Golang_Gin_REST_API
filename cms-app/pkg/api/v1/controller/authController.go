package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inder231/cms-app/pkg/api/v1/business"
	"github.com/inder231/cms-app/pkg/models"
)

// Signup handles user signup
// @Summary User Signup
// @Description Create a new user account with name, email, and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param {object} body models.UserCreationReqPayload true "Signup request body"
// @Success 201 {object} models.UserResponse
// @Failure 400
// @Failure 500
// @Router /signup [post]
func SignupUserController(c *gin.Context) {
	var user models.User

	// Parse JSON request body
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request body."})
		return
	}

	// Call business logic for user signup
	userResponse, err := business.SignupUserBusiness(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered!", "user": userResponse})
}

// Login represents the request body for signing up a new user
// @Summary User Login
// @Description Login request body
// @Tags Authentication
// @Accept json
// @Produce json
// @Param {object} body models.UserLoginReqPayload true "Login request body"
// @Success 200 {object} models.UserLoginResponse
// @Failure 400
// @Failure 500
// @Router /login [post]
func LoginUserController(c *gin.Context) {
	var loginPayload models.UserLoginReqPayload

	// Parse JSON request body
	err := c.ShouldBindJSON(&loginPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request body."})
		return
	}

	// Call business logic for user login
	response, err := business.LoginUserBusiness(loginPayload)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// VerifyTokenController handles token verification
// @Summary Verify Token
// @Description Verify the provided token and return the verification result
// @Tags Authentication
// @Accept json
// @Produce json
// @Param token query string true "Token to be verified"
// @Success 200 {object} map[string]string "Verification result"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 400 {object} map[string]string "Bad Request"
// @Router /verify-token [get]
func VerifyTokenController( c * gin.Context) {
	token := c.Query("token")
	msg, err := business.VerifyTokenBusiness(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
	}
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

// LogoutController handles user logout
// @Summary Logout a user
// @Description Invalidate the user's session token
// @Tags Authentication
// @Security JwtAuth
// @Success 200
// @Failure 500
// @Router /logout [post]
func LogoutController(c *gin.Context) {
    token := c.GetHeader("Authorization")
    if token == "" {
        c.JSON(http.StatusBadRequest, errors.New("missing auth token"))
        return
    }

    err := business.LogoutBusiness(token)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Logout successfull."})
}