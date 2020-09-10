package main

import (
	"Goless/GMP/apis"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/add_person", apis.AddPerson)

	router.GET("/get_person", apis.GetPerson)
	return router
}
