package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inder231/cms-app/pkg/api/v1/business"
	"github.com/inder231/cms-app/pkg/models"
)

// CreateBlogController handles the craetion of Blog with image
// @Summary Create Blog Post
// @Description Create a new blog post
// @Tags Blogs
// @Accept multipart/form-data
// @Produce json
// @param image formData file true "Blog Image"
// @Param blog body models.Blog true "Blog Details"
// @Success 201 {object} models.Blog
// @Failure 400
// @Failure 500
// @Router /blogs [post]
// @Security JwtAuth
func CreateBlogController(c *gin.Context) {
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

	var blog models.Blog
	if err := c.ShouldBind(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}

	// Call business logic
	response, err := business.CreateBlogBusiness(c, userId.(uint), file, blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetBlogsController retrieves all the blogs
// @Summary Get all blogs
// @Description Get Blogs
// @Tags Blogs
// @Produce json
// @Success 200 {array} models.Blog
// @Failure 500
// @Router /blogs [get]
// @Security JwtAuth
func GetBlogController(c *gin.Context) {
	blogs, err := business.ListBlogsBusiness()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve blogs"})
		return
	}
	c.JSON(http.StatusOK, blogs)
}

// DeleteBlogController handles the deletion of a blog post
// @Summary Delete a blog post
// @Description Deletes a blog post by its ID
// @Tags Blogs
// @Param id path string true "Blog ID"
// @Success 200
// @Failure 500
// @Router /blogs/{id} [delete]
// @Security JwtAuth
func DeleteBlogController(c *gin.Context) {
	id := c.Param("id")
	err := business.DeleteBlogBusiness(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}

// UpdateBlogController handles the updation of a blog post
// @Summary Update a blog post
// @Description Update the details of existing blog
// @Tags Blogs
// @Accept json
// @Produce json
// @Param id path string true "Blog ID"
// @Param blog body models.Blog true "Updated Blog Details"
// @Success 200 {object} models.Blog
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /blogs/{id} [put]
// @Security JwtAuth
func UpdateBlogController(c *gin.Context) {
	id := c.Param("id")

	var blog models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}

	updatedBlog, err := business.UpdateBlogBusiness(id, blog)
	if err != nil {
		if err.Error() == "blog not found" {
			c.JSON(http.StatusNotFound, gin.H{"message": "Blog not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, updatedBlog)
}
