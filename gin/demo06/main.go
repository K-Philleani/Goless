package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// 单文件上传
	router := gin.Default()

	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		c.String(200,  fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	router.Run(":9999")
}
