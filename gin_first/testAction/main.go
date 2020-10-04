package main

import (
	"Goless/gin_first/testAction/initDB"
	"Goless/gin_first/testAction/setInit"
)

func main() {
	initDB.Init()
	router := setInit.SetRouter()
	router.Run(":9000")

}
