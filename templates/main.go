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
	//iris.Templates().Delims("${", "}") this will change the behavior of {{.Property}} to ${.Property}
	//iris.Templates().Funcs(...)

	iris.Templates().Load("./tmpl/*.html", "mynamespace")

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Render("mypage.html", mypage{"My Page title", "Hello world!"})
	})

	//Get access to loaded (html/template) *template.Template with: iris.Templates().Templates
	iris.Listen(":8080")
}

/*
Want more render features?

Go to: https://github.com/iris-contrib/render

*/
