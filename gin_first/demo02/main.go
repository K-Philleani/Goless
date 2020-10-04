package main

import "github.com/gin-gonic/gin"

func main() {
	// 获取路径中的参数
	router := gin.Default()

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"name": name,
		})
	})

	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " 正在 " + action
		c.String(200, message)
	})

	router.Run(":9999")
}
