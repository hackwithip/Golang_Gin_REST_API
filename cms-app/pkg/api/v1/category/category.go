package category

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/inder231/cms-app/inits"
	"github.com/inder231/cms-app/pkg/models"
)

func CreateCategory(c *gin.Context) {
	var category models.Category

	userId, ok := c.Get("userId")

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authenticated"})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Image upload failed"})
		return
	}

	if err := c.ShouldBind(&category); err != nil {
		fmt.Println("---binding error---", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	category.CreatedBy = userId.(uint)

	categoryExist := inits.DB.Where("name = ?", category.Name).First(&category)

	if categoryExist.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Category with this name already exist.",
		})
		return
	}
	uploadDir := "uploads/category"
	// Construct the new filepath
	filePath := filepath.Join(uploadDir, file.Filename)

	// Save the uploaded file to the specified path
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save image"})
		return
	}
	category.CreatedBy = userId.(uint) 
	category.Image = filePath


	if result := inits.DB.Create(&category); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create category!"})
		return
	}

	categoryResp := models.CategoryCreationResponse{
		ID: category.ID,
		Name: category.Name,
		Image: category.Image,
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Category created", "category": categoryResp})
}

func ListCategories(c *gin.Context) {
	var categories []models.Category
	if result := inits.DB.Find(&categories); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve categories"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

// Delete Category
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	if result := inits.DB.Delete(&models.Category{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete category"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}
