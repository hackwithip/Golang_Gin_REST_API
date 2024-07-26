package models

import (
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null" json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	CategoryID  uint `json:"categoryId" form:"categoryId"`
	AuthorID    uint `json:"authorId" form:"authorId"`
	Category    Category `gorm:"foreignKey:CategoryID" form:"category" json:"category"`
	Author      Author   `gorm:"foreignKey:AuthorID" form:"author" json:"author"`
	Image       string `json:"image"`
	CreatedBy  uint 	`json:"createdBy"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
