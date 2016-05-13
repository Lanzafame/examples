package main

import (
	"github.com/kataras/iris"
)

func main() {

	iris.Config().Templates.Engine = iris.PongoEngine

	iris.Get("/", func(ctx *iris.Context) {

		err := ctx.Render("index.html", map[string]interface{}{"username": "iris", "is_admin": true})
		// OR
		//err := ctx.Render("index.html", pongo2.Context{"username": "iris", "is_admin": true})

		if err != nil {
			panic(err)
		}
	})

	println("Server is running at :8080")
	iris.Listen(":8080")
}
