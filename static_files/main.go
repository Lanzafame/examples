package main

import (
	"github.com/kataras/iris"
)

type page struct {
	Title string
}

func main() {
	iris.Config().Templates.Directory = "./templates/web/default"

	iris.Static("/css", "./resources/css", 1)
	iris.Static("/js", "./resources/js", 1)

	iris.Get("/", func(ctx *iris.Context) {
		err := ctx.Render("something.html", page{Title: "Home"})
		if err != nil {
			println(err.Error())
		}
	})

	println("Server is running at :8080")
	iris.Listen(":8080")
}
