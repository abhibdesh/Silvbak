package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	route := gin.Default()
	route.GET("/health-check", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	err := route.Run(":8080")
	if err != nil {
		panic(err)
	}
}
