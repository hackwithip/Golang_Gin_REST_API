package api

import (
	"github.com/gin-gonic/gin"
	"github.com/inder231/cms-app/pkg/api/middleware"
	"github.com/inder231/cms-app/pkg/api/v1/auth"
	"github.com/inder231/cms-app/pkg/api/v1/author"
	"github.com/inder231/cms-app/pkg/api/v1/blog"
	"github.com/inder231/cms-app/pkg/api/v1/category"
)

func RegisterRoutes( server *gin.Engine ){

    // Add more routes here...
    server.POST("/signup", auth.Signup)
    server.POST("/login", auth.Login)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)

	// Category routes
	authenticated.GET("/categories", category.ListCategories)
	authenticated.POST("/categories", category.CreateCategory)
	authenticated.DELETE("categories/:id", category.DeleteCategory)

	// Author routes
	authenticated.GET("/authors", author.ListAuthors)
	authenticated.POST("/authors", author.CreateAuthor)
	authenticated.DELETE("authors/:id", author.DeleteAuthor)

	// Blog routes
	authenticated.GET("/blogs", blog.GetBlog)
	authenticated.POST("/blogs", blog.CreateBlog)
	authenticated.PUT("/blogs/:id", blog.UpdateBlog)
	authenticated.DELETE("/blogs/:id", blog.DeleteBlog)

}