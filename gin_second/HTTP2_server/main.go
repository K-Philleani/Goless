package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
)

var html = template.Must(template.New("https").Parse(`
	<html>
	<head>
		<title>Https Test</title>
	<head>
	<h1 style="color:red;">Welcome, Ginner!</h1>
	</html>
`))

func main() {
	router := gin.Default()
	router.SetHTMLTemplate(html)

	router.GET("/", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			log.Println("Failed to push")
		}

		c.HTML(200, "https", gin.H{
			"status": "success",
		})
	})
	router.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
}