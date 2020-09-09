package main

import "GoLi/gin/action/hello/initRouter"

func main() {
	router := initRouter.SetupRouter()
	router.Run(":9999")
}