package service

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/inder231/cms-app/inits"
	"github.com/inder231/cms-app/pkg/models"
	"github.com/inder231/cms-app/pkg/services"
	"github.com/inder231/cms-app/pkg/utils"
)

func SignupUserService(user models.User) (models.UserResponse, error) {
	providedPassword := user.Password

	// Hash password
	hashedPassword, err := utils.HashedPassword(user.Password)
	if err != nil {
		return models.UserResponse{}, errors.New("failed to hash password")
	}
	user.Password = hashedPassword

	// Get Keycloak admin token
	keyCloakAdminUserName := os.Getenv("KEYCLOAK_ADMIN_USERNAME")
	keyCloakAdminPassword := os.Getenv("KEYCLOAK_ADMIN_PASSWORD")
	keyCloakAdminTokenResp, err := services.GetKeycloakAccessToken(keyCloakAdminUserName, keyCloakAdminPassword)
	if err != nil {
		return models.UserResponse{}, errors.New("failed to get Keycloak admin token")
	}

	accessToken := keyCloakAdminTokenResp.AccessToken

	// Create Keycloak user
	if !services.CreateKeycloakUser(accessToken, user.Email, providedPassword) {
		return models.UserResponse{}, errors.New("failed to create Keycloak user")
	}

	userId, err := services.GetKeycloakUser(accessToken, user.Email)
	if err != nil {
		return models.UserResponse{}, errors.New("failed to get Keycloak user")
	}

	// Store user in DB
	result := inits.DB.Create(&user)
	if result.Error != nil {
		return models.UserResponse{}, result.Error
	}

	// Store user in Keycloak table
	keycloakUser := models.KeycloakUser{
		ID:        userId,
		UserID:    user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  hashedPassword,
		RealmName: "master",
		Status:    "Inactive",
	}

	if err := inits.DB.Create(&keycloakUser).Error; err != nil {
		return models.UserResponse{}, errors.New("failed to create Keycloak user entry in database")
	}
	
	// Hold verification token in db
	verificationToken, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		return models.UserResponse{}, errors.New("failed to generate verification token")
	}

	token := models.Token{
        UserID:    user.ID,
        Token:     verificationToken,
        ExpiresAt: time.Now().Add(time.Minute * 15), // Expires in 15mins
    }
	if err := inits.DB.Create(&token).Error; err!= nil {
        return models.UserResponse{}, errors.New("failed to store verification token in database")
    }
	activationLink := fmt.Sprintf("http://localhost:3000/api/v1/verify-token?token=%s", verificationToken)
	
	emailBody := fmt.Sprintf(
		"Dear %s,\nThank you for registering with us! \n Below are the details provided for signup.\n\n Email: %s\n Password: %s\n\n Please given link to activate your account\n\n Link: %s. Best Regards!",
		user.Name, user.Email, providedPassword, activationLink,
	)
	if err := inits.TriggerEmailWorkflow(user.Email, "Welcome to Our Service", emailBody); err != nil {
		return models.UserResponse{}, errors.New("failed to send registration email")
	}

	return models.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func LoginUserService(loginPayload models.UserLoginReqPayload, user models.User) (models.UserLoginResponse, error) {
	// Validate password
	if !utils.CheckPasswordHash(loginPayload.Password, user.Password) {
		return models.UserLoginResponse{}, errors.New("invalid password")
	}

	// Get Keycloak token
	keyCloakAdminTokenResp, err := services.GetKeycloakAccessToken(loginPayload.Email, loginPayload.Password)
	if err != nil {
		return models.UserLoginResponse{}, errors.New("failed to get Keycloak admin token")
	}
	
	return models.UserLoginResponse{
		AccessToken: keyCloakAdminTokenResp.AccessToken, 
		RefreshToken: keyCloakAdminTokenResp.RefreshToken, 
		Message: "Login Successful.",
		}, nil
}

func VerifyTokenService(token string) (string, error) {
	var verification models.Token
    result := inits.DB.Where("token = ?", token).First(&verification)
    if result.Error != nil {
        return "", result.Error
    }

    if time.Now().After(verification.ExpiresAt) {
        return "", errors.New("token has expired")
    }

    var user models.User
	var keycloakuser models.KeycloakUser

    result = inits.DB.Where("id = ?", verification.UserID).First(&user)
    if result.Error != nil {
        return "", errors.New("failed to find user details")
    }

	keycloakUserFromDB := inits.DB.Where("user_id = ?", verification.UserID).First(&keycloakuser)
	if keycloakUserFromDB.Error != nil {
        return "", errors.New("failed to find keycloak user details")
    }

    // Update user status to active
    user.Status = "Active" // Assuming you have a `Status` field in the `User` model
	keycloakuser.Status = "Active"
    result = inits.DB.Save(&user)
	saveKeyCloakUser := inits.DB.Save(&keycloakuser)

	if saveKeyCloakUser.Error!= nil {
        return "", errors.New("failed to update keycloak user details, try again later")
    }
    if result.Error != nil {
        return "", errors.New("failed to update user details, try again later")
    }

    // Optionally, delete the verification token
    result = inits.DB.Delete(&verification)
    if result.Error != nil {
        return "", errors.New("internal server error")
    }

    return "User validated successfully.", nil
}

func LogoutUserService(token string) error {
	    // Call Keycloak to invalidate the token
		err := services.LogoutKeycloakUser(token)
		if err != nil {
			return fmt.Errorf("failed to logout user: %w", err)
		}
		return nil
}