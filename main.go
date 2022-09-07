package main

import (
	"belajar-gin/explore/views"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", views.Index)
	router.Run("localhost:7000")
}
