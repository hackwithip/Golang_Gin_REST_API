package models

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	ID uint `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	UserID uint `json:"userId"`
	Token string `json:"token"`
	ExpiresAt time.Time `json:"expiresAt" gorm:"not null"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

type TokenValidationResponse struct {
	Message string `json:"message"`
}