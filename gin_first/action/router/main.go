package main
import (
	"GoLi/gin/action/router/setRouter"
)
func main() {
	router := setRouter.SetRouter()
	router.Run(":9999")
}
