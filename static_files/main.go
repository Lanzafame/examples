package main

import (
	"github.com/kataras/iris"
)

type page struct {
	Title string
}

func main() {
	//optionally
	iris.TemplateDelims("${", "}")
	//
	iris.Templates("./templates/*.html")

	iris.Static("/css", "./resources/css", 1)
	iris.Static("/js", "./resources/js", 1)

	iris.Get("/", func(c *iris.Context) {
		err := c.RenderFile("something.html", page{Title: "Home"})
		if err != nil {
			println(err.Error())
		}
	})

	println("Server is running at :8080")
	iris.Listen()
}
