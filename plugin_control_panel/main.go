package main

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/plugin/iriscontrol"
)

func main() {

	iris.Plugins().Add(iriscontrol.Web(9090, map[string]string{
		"irisusername1": "irispassword1",
		"irisusername2": "irispassowrd2",
	}))

	iris.Get("/", func(ctx *iris.Context) {
	})

	iris.Post("/something", func(ctx *iris.Context) {
	})

	fmt.Printf("Iris is listening on :%d", 8080)
	iris.Listen(":8080")
}
