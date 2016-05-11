package main

import (
	"github.com/kataras/iris"
)

func main() {
	iris.Config().Render.Layout = "layouts/layout"

	iris.Get("/", func(ctx *iris.Context) {
		if err := ctx.Render("page1", nil); err != nil {
			panic(err)
		}
	})

	iris.Listen(":80")
}
