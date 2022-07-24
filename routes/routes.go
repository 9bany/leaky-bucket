package routes

import (
	"9bany/leaky-bucket/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/hello", middleware.Middleware, sayHello)
}

func sayHello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"messgae": "Hello",
	})
}
