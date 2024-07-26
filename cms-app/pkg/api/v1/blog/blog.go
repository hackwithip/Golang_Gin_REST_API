package blog

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/inder231/cms-app/inits"
	"github.com/inder231/cms-app/pkg/models"
)

func CreateBlog(c *gin.Context) {
	var blog models.Blog

	userId, ok := c.Get("userId")

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authenticated"})
		return
	}
	
	if err := c.ShouldBind(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}

	// Check if author already exist with same name
	blogExist := inits.DB.Where("title = ?", blog.Title).First(&blog)

	if blogExist.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Blog with this title already exists.",
		})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Image upload failed"})
		return
	}

	uploadDir := "uploads/blog"
	// Construct the new filepath
	filePath := filepath.Join(uploadDir, file.Filename)

	// Save the uploaded file to the specified path
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save image"})
		return
	}
	blog.CreatedBy = userId.(uint) 
	blog.Image = filePath


	result := inits.DB.Create(&blog)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": result.Error})
		return
	}
	c.JSON(http.StatusCreated, blog)
}

func GetBlog(c *gin.Context) {
	var blog []models.Blog
	result := inits.DB.Find(&blog)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create blog"})
		return
	}
	c.JSON(http.StatusOK, blog)

}

func DeleteBlog(c *gin.Context) {
	id := c.Param("id")
	result := inits.DB.Delete(&models.Blog{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete blog"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}

func UpdateBlog(c *gin.Context) {
	var blog models.Blog
	id := c.Param("id")

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}
	var existingBlog models.Blog
	if err := inits.DB.First(&existingBlog, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Blog not found"})
		return
	}
	existingBlog.Title = blog.Title
	existingBlog.Description = blog.Description
	existingBlog.AuthorID = blog.AuthorID
	existingBlog.CategoryID = blog.CategoryID
	existingBlog.Image = blog.Image
	if err := inits.DB.Save(&existingBlog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update blog"})
		return
	}

	c.JSON(http.StatusOK, existingBlog)
}
