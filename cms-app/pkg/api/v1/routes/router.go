package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/inder231/cms-app/configs"
	"github.com/inder231/cms-app/pkg/api/middleware"
	"github.com/inder231/cms-app/pkg/api/v1/controller"
)


func RegisterRoutes( server *gin.Engine) {

	// Authentication Routes
	server.POST(configs.SignupRoute, controller.SignupUserController)
	server.POST(configs.LoginRoute, controller.LoginUserController)
	server.GET(configs.VerifyToken, controller.VerifyTokenController)
	// Logout Route
	server.POST(configs.LogoutRoute, controller.LogoutController)
	
	// Auth Middleware
	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	

	// Category Routes
	authenticated.GET(configs.GetCategoriesRoute, controller.GetCategoryListController)
	authenticated.POST(configs.CreateCategoriesRoute, controller.CreateCategoryController)
	authenticated.DELETE(configs.DeleteCategoryRoute, controller.DeleteCategoryController)

	// Author Routes
	authenticated.GET(configs.GetAuthorsRoute, controller.GetAuthorListController)
	authenticated.POST(configs.CreateAuthorRoute, controller.CreateAuthorController)
	authenticated.DELETE(configs.DeleteAuthorsRoute, controller.DeleteAuthorController)

	// Blog Routes
	authenticated.GET(configs.GetBlogsRoute, controller.GetBlogController)
    authenticated.POST(configs.CreateBlogsRoute, controller.CreateBlogController)
    authenticated.PUT(configs.UpdateBlogRoute, controller.UpdateBlogController)
    authenticated.DELETE(configs.DeleteBlogRoute, controller.DeleteBlogController)

}