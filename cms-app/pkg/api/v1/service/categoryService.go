package service

import (
	"errors"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/inder231/cms-app/inits"
	"github.com/inder231/cms-app/pkg/models"
)

func CreateCategoryService(c *gin.Context, userId uint, file *multipart.FileHeader, category models.Category) (models.CategoryCreationResponse, error) {
	// Handle file upload
	uploadDir := "uploads/category"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return models.CategoryCreationResponse{}, errors.New("failed to create directory for uploads")
	}
	
	// Construct the new filepath
	filePath := filepath.Join(uploadDir, file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		return models.CategoryCreationResponse{}, errors.New("failed to save image")
	}

	// Set the file path and created by fields
	category.Image = filePath
	category.CreatedBy = userId

	// Save category in the database
	if result := inits.DB.Create(&category); result.Error != nil {
		return models.CategoryCreationResponse{}, errors.New("failed to create category")
	}

	// Prepare the response
	categoryResp := models.CategoryCreationResponse{
		ID:    category.ID,
		Name:  category.Name,
		Image: category.Image,
	}

	return categoryResp, nil

}

func ListCategoriesService() ([]models.Category, error) {
	var categories []models.Category
	if result := inits.DB.Find(&categories); result.Error != nil {
		return nil, errors.New("failed to retrieve categories")
	}
	return categories, nil
}

func DeleteCategoryService(id string) error {
	if result := inits.DB.Delete(&models.Category{}, id); result.Error != nil {
		return errors.New("failed to delete category")
	}
	return nil
}
