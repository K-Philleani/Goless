package main

import (
	"Goless/go_mysql/apis"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/add_person", apis.AddPerson)
	router.GET("/get_person", apis.GetPerson)
	router.GET("/get_person_one", apis.GetPersonOne)
	return router
}
