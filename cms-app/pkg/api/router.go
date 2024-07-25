package api

import (
	"github.com/gin-gonic/gin"
	"github.com/inder231/cms-app/pkg/api/v1/auth"
)

func RegisterRoutes( server *gin.Engine ){

    // Add more routes here...
    server.POST("/signup", auth.Signup)
    server.POST("/login", auth.Login)

}