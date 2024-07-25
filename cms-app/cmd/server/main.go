package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inder231/cms-app/pkg/utils"
)

func main() {

	// Load environment variables
	utils.LoadEnv()

    router := gin.Default()

    router.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
    })
	
	// Start the server
    router.Run()

}
