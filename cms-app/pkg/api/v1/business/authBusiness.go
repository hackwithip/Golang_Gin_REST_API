package business

import (
	"errors"

	"github.com/inder231/cms-app/inits"
	"github.com/inder231/cms-app/pkg/api/v1/service"
	"github.com/inder231/cms-app/pkg/models"
	"gorm.io/gorm"
)

func SignupUserBusiness(user models.User) (models.UserResponse, error) {
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return models.UserResponse{}, errors.New("name, email, and password are required")
	}
	// Check if email already exists
	var existingUser models.User
	isUserAlreadyExist := inits.DB.Where("email = ?", user.Email).Find(&existingUser)

	if isUserAlreadyExist.RowsAffected > 0 {
		return models.UserResponse{}, errors.New("user with this email already exists")
	}
	return service.SignupUserService(user)
}

func LoginUserBusiness(loginPayload models.UserLoginReqPayload) (models.UserLoginResponse, error) {
	if loginPayload.Email == "" || loginPayload.Password == "" {
		return models.UserLoginResponse{}, errors.New("email and password are required")
	}

	var existingUser models.User
	result := inits.DB.Where("email = ?", loginPayload.Email).First(&existingUser)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return models.UserLoginResponse{}, errors.New("user not registered, please signup")
		}
		return models.UserLoginResponse{}, errors.New("failed to check user existence")
	}

	if existingUser.Status != "Active" {
		return models.UserLoginResponse{}, errors.New("user is not active, please verify your email")
	}

	// User from keycloak table
	var keycloakUser models.KeycloakUser
    result = inits.DB.Where("email = ?", loginPayload.Email).First(&keycloakUser)
    if result.Error!= nil {
        if result.Error == gorm.ErrRecordNotFound {
            return models.UserLoginResponse{}, errors.New("user not registered, please signup")
        }
        return models.UserLoginResponse{}, errors.New("failed to check user existence in keycloak table")
    }

    if keycloakUser.Status!= "Active" {
        return models.UserLoginResponse{}, errors.New("user in keycloak is not active, please verify your email")
    }

	return service.LoginUserService(loginPayload, existingUser)
}

func VerifyTokenBusiness(token string) (string, error) {
	if token == "" {
		return "", errors.New("token is required")
	}
	return service.VerifyTokenService(token)
}

func LogoutBusiness(token string) error {
	return service.LogoutUserService(token)
}