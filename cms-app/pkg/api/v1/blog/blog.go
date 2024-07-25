package blog

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inder231/cms-app/inits"
	"github.com/inder231/cms-app/pkg/models"
)

func CreateBlog(c *gin.Context) {
	var blog models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}
	result := inits.DB.Create(&blog)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create blog"})
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
