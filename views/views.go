package views

import (
	// "fmt"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	// data := map[string]interface{}{
	// 	"request": "test",
	// }
	c.IndentedJSON(200, gin.H{"message": "hello world", "data": c.Request.Header})
}
