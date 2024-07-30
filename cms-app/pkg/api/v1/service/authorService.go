package service

import (
	"errors"
	"mime/multipart"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/inder231/cms-app/inits"
	"github.com/inder231/cms-app/pkg/models"
)

func CreateAuthorService(c *gin.Context, userId uint, file *multipart.FileHeader, author models.Author) (models.AuthorCreationResponse, error) {

	// Save the file and update the author details
	uploadDir := "uploads/author"
	filePath := filepath.Join(uploadDir, file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		return models.AuthorCreationResponse{}, errors.New("failed to save image")
	}

	author.CreatedBy = userId
	author.Image = filePath

	// Save author in the database
	if err := inits.DB.Create(&author).Error; err != nil {
		return models.AuthorCreationResponse{}, errors.New("failed to create category")
	}
	authorResp := models.AuthorCreationResponse{
		ID:    author.ID,
		Name:  author.Name,
		Image: author.Image,
	}
	return authorResp, nil
}

func ListAuthorsService() ([]models.Author, error){
	var authors []models.Author
	if result := inits.DB.Find(&authors); result.Error != nil {
		return nil, errors.New("failed to retrieve authors")
	}
	return authors, nil
}

func DeleteAuthorService(id string) error {
	if result := inits.DB.Delete(&models.Author{}, id); result.Error != nil {
		return errors.New("failed to delete author")
	}
	return nil
}
