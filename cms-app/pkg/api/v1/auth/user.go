package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inder231/cms-app/inits"
	"github.com/inder231/cms-app/pkg/models"
	"github.com/inder231/cms-app/pkg/utils"
	"gorm.io/gorm"
)

func Signup(c *gin.Context) {

	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not paser request body."})
		return
	}

	// Check if email already exist
	isUserAlreadyExist := inits.DB.Where("email = ?", user.Email).Find(&user)

	if isUserAlreadyExist.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User with this email already exists.",
		})
		return
	}

	// Hash password
	hashedPassword, err := utils.HashedPassword(user.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Update User password with hashed password
	user.Password = hashedPassword

	// Store user in DB
	result := inits.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	err = inits.TriggerEmailWorkflow(user.Email, "Welcome to Our Service", "Thank you for registering with us!")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to send registration email"})
		return
	}
	userResponse := models.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered!", "user": userResponse})
}

func Login(c *gin.Context) {
	// Use two variables one to store request body user and one to hold user from db
	var user models.User
	var existingUser models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not paser request body."})
		return
	}

	// Check if email already exist
	result := inits.DB.Where("email = ?", user.Email).First(&existingUser)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User not registered, please signup!"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check user existence"})
		}
		return
	}
	// Validate password
	passwordIsValid := utils.CheckPasswordHash(user.Password, existingUser.Password)

	if !passwordIsValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password."})
		return
	}

	// Generate token
	token, err := utils.GenerateToken(existingUser.Email, existingUser.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})

}
