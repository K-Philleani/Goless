package main

import (
	"Goless/gin/testAction/initDB"
	"Goless/gin/testAction/setInit"
)

func main() {
	initDB.Init()
	router := setInit.SetRouter()
	router.Run(":9000")

}
