package views

import (
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	data := map[string]interface{}{
		"request": c.Header,
	}
	c.IndentedJSON(200, gin.H{"message": "hello world", "data": data})
}
