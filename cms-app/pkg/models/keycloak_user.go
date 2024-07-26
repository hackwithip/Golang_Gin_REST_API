package models

import (
	"time"

	"gorm.io/gorm"
)

type KeycloakUser struct {
	ID       string 		 `gorm:"primaryKey" json:"id"`
	UserID   uint   		 `json:"userId"`
	User     User            `gorm:"foreignKey:ID" json:"user"`
	Name     string  		 `gorm:"size;255;not null" json:"name"`
	Email    string 		 `gorm:"size:255;not null;unique" json:"email"`
	Password  string         `gorm:"size:255;not null" json:"password"`
	RealmName string 		 `json:"realmName"`
	Status    string         `json:"status"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

