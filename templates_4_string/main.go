package main

import (
	"github.com/kataras/iris"
)

func main() {
	iris.Config().Render.Template.Layout = "layouts/layout.html" // default ""
	iris.Config().Render.Template.Minify = true                  // default true already
	iris.Get("/", func(ctx *iris.Context) {
		if str, err := ctx.RenderString("page1.html", nil); err != nil {
			panic(err)
		} else {
			println(str)
			ctx.Write("The plain content of the template is: %s", str)
		}

	})

	// you can do the same without context also using iris.Templates() :
	if str, err := iris.Templates().RenderString("page1.html", nil); err != nil {
		panic(err)
	} else {
		println(str)
	}

	//

	println("Server is running at: 8080")
	iris.Listen(":8080")
}
