package service

import (
	"errors"
	"mime/multipart"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/inder231/cms-app/inits"
	"github.com/inder231/cms-app/pkg/models"
)

func CreateBlogService(c *gin.Context, userId uint, file *multipart.FileHeader, blog models.Blog) (models.Blog, error) {
	// Save the file and update the blog details
	uploadDir := "uploads/blog"
	filePath := filepath.Join(uploadDir, file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		return models.Blog{}, errors.New("failed to save image")
	}

	blog.CreatedBy = userId
	blog.Image = filePath

	// Save blog in the database
	if err := inits.DB.Create(&blog).Error; err != nil {
		return models.Blog{}, errors.New("failed to create blog")
	}

	return blog, nil
}

func ListBlogsService() ([]models.Blog, error) {
	var blogs []models.Blog
	result := inits.DB.Find(&blogs)
	if result.Error != nil {
		return nil, errors.New("failed to retrieve blogs")
	}
	return blogs, nil
}

func DeleteBlogService(id string) error {
	result := inits.DB.Delete(&models.Blog{}, id)
	if result.Error != nil {
		return errors.New("failed to delete blog")
	}
	return nil
}

func UpdateBlogService(id string, updatedBlog models.Blog) (models.Blog, error) {
	var existingBlog models.Blog
	if err := inits.DB.First(&existingBlog, id).Error; err != nil {
		return models.Blog{}, errors.New("blog not found")
	}

	// Update fields
	existingBlog.Title = updatedBlog.Title
	existingBlog.Description = updatedBlog.Description
	existingBlog.AuthorID = updatedBlog.AuthorID
	existingBlog.CategoryID = updatedBlog.CategoryID
	existingBlog.Image = updatedBlog.Image

	if err := inits.DB.Save(&existingBlog).Error; err != nil {
		return models.Blog{}, errors.New("failed to update blog")
	}

	return existingBlog, nil
}
