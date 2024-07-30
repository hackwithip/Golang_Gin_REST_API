package business

import (
	"errors"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/inder231/cms-app/inits"
	"github.com/inder231/cms-app/pkg/api/v1/service"
	"github.com/inder231/cms-app/pkg/models"
)

func CreateCategoryBusiness(c *gin.Context, userId uint, file *multipart.FileHeader, category models.Category) (models.CategoryCreationResponse, error) {
	if category.Name == "" {
		return models.CategoryCreationResponse{}, errors.New("category name is required")
	}

	// Check if category name already exists
	var existingCategory models.Category
	categoryExist := inits.DB.Where("name = ?", category.Name).First(&existingCategory)
	if categoryExist.RowsAffected > 0 {
		return models.CategoryCreationResponse{}, errors.New("category with this name already exists")
	}

	return service.CreateCategoryService(c, userId, file, category)
}

func ListCategoriesBusiness() ([]models.Category, error) {
	return service.ListCategoriesService()
}

func DeleteCategoryBusiness(id string) error {
	return service.DeleteCategoryService(id)
}