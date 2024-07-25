package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inder231/cms-app/inits"
)

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
	
	// Start the server
    router.Run()

}
