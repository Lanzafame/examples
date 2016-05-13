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
	//iris.Config().Templates.Standar.Left = "${"  // Default is "{{"
	//iris.Config().Templates.Standar.Right = "}" // Default is "}}"
	//iris.Config().Templates.Standar.Funcs = template.FuncMap(...)

	//iris.Config().Templates.Directory = "templates" // Default is "templates"
	iris.Config().Templates.IsDevelopment = true // rebuild the templates on each refresh. Default is false
	//iris.Config().Templates.Layout = "layout.html" // means: ./templates/layout.html.  Default is ""
	iris.Config().Templates.Gzip = true // Default is false

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Render("mypage.html", mypage{"My Page title", "Hello world!"}) //, "otherLayout" <- to override the layout
	})

	println("Server is running at :8080")
	iris.Listen(":8080")
}
