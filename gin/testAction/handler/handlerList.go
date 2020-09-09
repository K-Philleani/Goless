package handler

import (
	"Goless/gin/testAction/model"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func ApiList (router *gin.Engine) {
	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("./../templates/*")
	} else {
		router.LoadHTMLGlob("templates/*")
	}

	// 路由未分组
	router.GET("/user/:username", func(c *gin.Context) {
		username := c.Param("username")
		c.String(200, "Hello" + username)
	})

	router.GET("/username", func(c *gin.Context) {
		username := c.DefaultQuery("username", "lisan")
		age := c.Query("age")
		c.String(200, "hello: %s/%s", username, age)
	} )

	// 路由分组
	index := router.Group("/")
	{
		index.GET("", retGet)
	}

	temRouter := router.Group("/tem")
	{
		temRouter.GET("", retTem)
		temRouter.POST("/register", retRegister)
	}

}

func retGet(c *gin.Context) {
	c.String(200, "测试用例：GET, /")
}

func retTem(c *gin.Context) {
	c.HTML(200, "index.tmpl", gin.H{
		"title": "Hello Gin " + strings.ToLower(c.Request.Method) + " method",
	})
}

func retRegister(c *gin.Context) {
	var temUser model.TemModel
	if err := c.ShouldBind(&temUser); err != nil {
		log.Println("err:", err)
		return
	}
	println("email: ", temUser.Email, "password: ", temUser.Password, "passwordAgain: ", temUser.PasswordAgain)


	//if err := c.ShouldBind(&temUser); err != nil {
	//	log.Println("err ->", err.Error())
	//	c.String(http.StatusBadRequest, "输入的数据不合法")
	//} else {
	//	println("email: ", temUser.Email, "password: ", temUser.Password, "passwordAgain: ", temUser.PasswordAgain)
	//	c.Redirect(http.StatusMovedPermanently, "/")
	//}
}


