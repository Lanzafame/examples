package main

import (
	"github.com/kataras/iris"
)

type mypage struct {
	Title   string
	Message string
}

func main() {

	//optionally - before the load.
	iris.Config().Render.Delims.Left = "${" // Default is "{{"
	iris.Config().Render.Delims.Right = "}" //this will change the behavior of {{.Property}} to ${.Property}. Default is "}}"
	//iris.Config().Render.Funcs = template.FuncMap(...)

	//iris.Config().Render.Directory = "templates" // Default is "templates"

	iris.Config().Render.Layout = "layout" // Default is ""
	iris.Config().Render.Gzip = true       // Default is false
	iris.Get("/", func(ctx *iris.Context) {
		ctx.Render("mypage", mypage{"My Page title", "Hello world!"}) //, "otherLayout" <- to override the layout
	})

	println("Server is running at :8080")
	iris.Listen(":8080")
}
