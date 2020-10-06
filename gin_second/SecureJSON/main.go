package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/someJson", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		c.SecureJSON(200, names)
	})
	router.Run(":8080")
}
