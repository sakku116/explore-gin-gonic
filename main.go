package main

import (
	"belajar-gin/explore/views"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	fmt.Println("test")
	router.GET("/", views.Index)
	router.Run("localhost:7000")
}
