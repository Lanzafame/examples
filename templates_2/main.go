package main

import (
	"github.com/kataras/iris"
)

func main() {
	iris.Config().Templates.Layout = "layouts/layout.html"
	iris.Get("/", func(ctx *iris.Context) {
		if err := ctx.Render("page1.html", nil); err != nil {
			panic(err)
		}
	})

	println("Server is running at: 8080")
	iris.Listen(":8080")
}
