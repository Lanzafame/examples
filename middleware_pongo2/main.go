package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/pongo2"
)

func main() {
	iris.Use(pongo2.Pongo2("./templates")) // or .Pongo2() defaults to "./"

	iris.Config().Render.Directory = "change_this_because_its_defaults_to_templates_directory_also"

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Set("template", "index.html")
		ctx.Set("data", map[string]interface{}{"username": "iris", "is_admin": true})
	})

	println("Server is running at :8080")
	iris.Listen(":8080")
}
