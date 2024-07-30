package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inder231/cms-app/pkg/api/v1/business"
	"github.com/inder231/cms-app/pkg/models"
)

// CreateCategoryController handles the creation of a category
// @Summary Create a category
// @Description Creates a new category with an image
// @Tags Category
// @Accept multipart/form-data
// @Produce json
// @Param image formData file true "Category Image"
// @Param category body models.Category true "Category Details"
// @Success 201 {object} models.CategoryCreationResponse
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /category [post]
// @Security JwtAuth
func CreateCategoryController(c *gin.Context) {
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

	var category models.Category
	if err := c.ShouldBind(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	// Call business logic for category creation
	categoryResp, err := business.CreateCategoryBusiness(c, userId.(uint), file, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Category created", "category": categoryResp})
}

// ListCategoriesController retrieves and returns all categories
// @Summary List all categories
// @Description Fetches a list of all categories from the database
// @Tags Category
// @Produce json
// @Success 200 {array} models.Category
// @Failure 500
// @Router /categories [get]
// @Security JwtAuth
func GetCategoryListController(c *gin.Context) {
	// Call business logic to list categories
	categories, err := business.ListCategoriesBusiness()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve categories"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

// DeleteCategoryController handles the deletion of a category
// @Summary Delete a category
// @Description Deletes a category by its ID
// @Tags Category
// @Param id path string true "Category ID"
// @Success 200
// @Failure 500
// @Router /category/{id} [delete]
// @Security JwtAuth
func DeleteCategoryController(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID provided"})
        return
	}
	err := business.DeleteCategoryBusiness(id)
	if err!= nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete category"})
        return
    }
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}