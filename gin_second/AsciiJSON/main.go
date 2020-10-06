package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 使用 AsciiJSON 生成具有转义的非 ASCII 字符的 ASCII-only JSON
func main() {
	router := gin.Default()
	router.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "Go语言",
			"tag": "<br>",
		}
		c.AsciiJSON(http.StatusOK, data)
	})

	router.Run(":8080")
}
