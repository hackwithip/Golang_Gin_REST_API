package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/inder231/cms-app/cmd/docs"
	"github.com/inder231/cms-app/inits"

	"github.com/inder231/cms-app/pkg/api/v1/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Documenting API
// @version 1
// Description Sample Description

// @contact.name Inder
// @contact.url http://github.com/inder231

// @host localhost:3000
// @BasePath /api/v1
// @securityDefinitions.apiKey JwtAuth
// @in header
// @name Authorization
// @type apiKey

func main() {

	// Load environment variables
	inits.LoadEnv()
    // Initialize Postgres DB
    inits.InitPgDB()

    router := gin.Default()

    router.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
    })
    // Swagger setup
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    routes.RegisterRoutes(router)
	
	// Start the server
    router.Run()

}
 