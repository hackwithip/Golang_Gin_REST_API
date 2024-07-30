package business

import (
	"errors"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/inder231/cms-app/inits"
	"github.com/inder231/cms-app/pkg/api/v1/service"
	"github.com/inder231/cms-app/pkg/models"
)

func CreateAuthorBusiness(c *gin.Context, userId uint, file *multipart.FileHeader, author models.Author) (models.AuthorCreationResponse, error) {
	if author.Name == "" {
		return models.AuthorCreationResponse{}, errors.New("author name is required")
	}

	// Check if author already exists with the same name
	authorExist := inits.DB.Where("name = ?", author.Name).First(&author)
	
	if authorExist.RowsAffected > 0 {
		return models.AuthorCreationResponse{}, errors.New("category with this name already exists")
	}

	return service.CreateAuthorService(c, userId, file, author)
}

func ListAuthorsBusiness() ([]models.Author, error) {
    return service.ListAuthorsService()
}

func DeleteAuthorBusiness(id string) error {
    return service.DeleteAuthorService(id)
}