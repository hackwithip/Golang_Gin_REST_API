package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inder231/cms-app/pkg/api/v1/business"
	"github.com/inder231/cms-app/pkg/models"
)

// CreateAuthor creates a new author
// @Summary Create a new author
// @Description Create a new author with the provided details and image
// @Tags Authors
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "Author's name"
// @Param image formData file true "Author's image"
// @Success 201 {object} models.AuthorCreationResponse
// @Failure 400
// @Failure 401
// @Failure 500
// @Security JwtAuth
// @Router /authors [post]
func CreateAuthorController(c *gin.Context) {
		// Extract user ID from context
		userId, ok := c.Get("userId")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authenticated"})
			return
		}
	
		// Extract file and form data
		file, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Image upload failed"})
			return
		}
	
		var author models.Author
		if err := c.ShouldBind(&author); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
			return
		}
	
		// Call business logic
		response, err := business.CreateAuthorBusiness(c, userId.(uint), file, author)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	
		c.JSON(http.StatusCreated, gin.H{"message": "Author created successfully!", "author": response})
}

// ListAuthors returns a list of all authors from the database
// @Summary List all authors
// @Description Returns a list of all authors from the database
// @Tags Authors
// @Produce json
// @Success 200
// @Failure 500
// @Security JwtAuth
// @Router /authors [get]
func GetAuthorListController(c *gin.Context) {
	authors, err := business.ListAuthorsBusiness()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve categories"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"categories": authors})
}

// DeleteAuthor deletes an author by ID
// @Summary Delete an author
// @Description Deletes an author from the database by their ID
// @Tags Authors
// @Produce json
// @Param id path string true "Author ID"
// @Success 200
// @Failure 500
// @Security JwtAuth
// @Router /authors/{id} [delete]
func DeleteAuthorController(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID provided"})
        return
	}
	err := business.DeleteAuthorBusiness(id)
	if err!= nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete author"})
        return
    }
	c.JSON(http.StatusOK, gin.H{"message": "Author deleted successfully"})
}