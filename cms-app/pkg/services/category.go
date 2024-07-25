package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inder231/cms-app/inits"
	"github.com/inder231/cms-app/pkg/models"
)

func createCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	if result := inits.DB.Create(&category); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Name must be unique"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Category created", "category": category})
}

func listCategories(c *gin.Context) {
	var categories []models.Category
	if result := inits.DB.Find(&categories); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve categories"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

// Delete Category
func deleteCategory(c *gin.Context) {
	id := c.Param("id")
	if result := inits.DB.Delete(&models.Category{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete category"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}
