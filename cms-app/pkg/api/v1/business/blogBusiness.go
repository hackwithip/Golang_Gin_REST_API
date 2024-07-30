package business

import (
	"errors"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/inder231/cms-app/inits"
	"github.com/inder231/cms-app/pkg/api/v1/service"
	"github.com/inder231/cms-app/pkg/models"
)

func CreateBlogBusiness(c *gin.Context, userId uint, file *multipart.FileHeader, blog models.Blog) (models.Blog, error) {
	// Validate blog data
	if blog.Title == "" {
		return models.Blog{}, errors.New("blog title is required")
	}

	// Check if blog with the same title already exists
	blogExist := inits.DB.Where("title = ?", blog.Title).First(&blog)
	if blogExist.RowsAffected > 0 {
		return models.Blog{}, errors.New("blog with this title already exists")
	}

	return service.CreateBlogService(c, userId, file, blog)
}

func ListBlogsBusiness() ([]models.Blog, error) {
	return service.ListBlogsService()
}

func DeleteBlogBusiness(id string) error {
	return service.DeleteBlogService(id)
}

func UpdateBlogBusiness(id string, blog models.Blog) (models.Blog, error) {
	// Validate blog data if necessary
	if blog.Title == "" {
		return models.Blog{}, errors.New("blog title is required")
	}

	return service.UpdateBlogService(id, blog)
}
