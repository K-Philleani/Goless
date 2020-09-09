package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 获取GET参数
	router := gin.Default()

	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.Query("firstname")
		lastname := c.Query("lastname")
		message := "Hello" + firstname + lastname
		c.String(http.StatusOK, message)
	})
	router.Run(":9999")
}
