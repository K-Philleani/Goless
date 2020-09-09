package setInit

import (
	"Goless/gin/testAction/handler"
	"github.com/gin-gonic/gin"
)


func SetRouter() *gin.Engine {
	router := gin.Default()

	router.Static("/statics","./statics")
	handler.ApiList(router)
	return router
}
