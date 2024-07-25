package author

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inder231/cms-app/inits"
	"github.com/inder231/cms-app/pkg/models"
)

func CreateAuthor(c *gin.Context) {
	var author models.Author

	userId, ok := c.Get("userId")

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authenticated"})
		return
	}
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}

	author.CreatedBy = userId.(uint) 

	// Check if author already exist with same name
	authorExist := inits.DB.Where("name = ?", author.Name).First(&author)

	if authorExist.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Author with this name already exists.",
		})
		return
	}

	result := inits.DB.Create(&author)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create author!"})
		return
	}

	authorResp := models.AuthorCreationResponse{
		ID:    author.ID,
        Name:  author.Name,
        Image: author.Image,
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Author created successfully!", "author": authorResp})
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
