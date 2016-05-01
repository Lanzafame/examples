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
	//iris.Config().Render.Delims = iris.Delims{Left:"${", Right: "}"} this will change the behavior of {{.Property}} to ${.Property}
	//iris.Config().Render.Funcs = template.FuncMap(...)

	//iris.Config().Render.Directory = "tmpl"
	iris.Config().Render.Layout = "layout"

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Render("mypage", mypage{"My Page title", "Hello world!"}) //, iris.HTMLOptions{"otherLayout"}) <- to override
	})

	println("Server is running at :8080")
	iris.Listen(":8080")
}
