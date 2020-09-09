package setRouter

import "github.com/gin-gonic/gin"

func SetRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(200, "name: %s", name)
	})

	return router
}
