package models

import (
	"time"

	"gorm.io/gorm"
)

// UserResponse represents the response data after signup
// @Description Response data after user signup
type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserCreationReqPayload respresents the user creation payoad for signup
// @Description Request payload for user signup
type UserCreationReqPayload struct {
	Name  string `json:"name" example:"Inder"`
	Email string `json:"email" example:"inderp@moneymul.com"`
	Password string `json:"password" example:"inder@123"`
}

// UserLoginReqPayload represetnts the user login payload
// @Description Reqest payload for user login
type UserLoginReqPayload struct {
	Email    string `json:"email" example:"inderp@moneymul.com"`
	Password string `json:"password" example:"inder@123"`
}

// UserLoginResponse represents user login response
type UserLoginResponse struct {
	Message string `json:"message"`
	AccessToken string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
type User struct {
	ID        uint           `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Name      string         `json:"name" gorm:"size:255;not null"`
	Email     string         `json:"email" gorm:"size:255;not null;unique"`
	Password  string         `json:"password" gorm:"size:255;not null"`
	Status    string		 `json:"status" gorm:"default:'Inactive'"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}