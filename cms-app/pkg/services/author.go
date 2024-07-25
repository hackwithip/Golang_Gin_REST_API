package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inder231/cms-app/inits"
	"github.com/inder231/cms-app/pkg/models"
)

func CreateAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}
	result := inits.DB.Create(&author)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create author"})
		return
	}
	c.JSON(http.StatusCreated, author)
}

// List all authors
func ListAuthors(c *gin.Context) {
	var authors []models.Author
	result := inits.DB.Find(&authors)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch authors"})
		return
	}
	c.JSON(http.StatusOK, authors)
}

// Delete an author
func DeleteAuthor(c *gin.Context) {
	id := c.Param("id")
	result := inits.DB.Delete(&models.Author{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete author"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Author deleted successfully"})
}
