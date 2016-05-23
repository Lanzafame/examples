package main

import (
	"fmt"
	"time"

	"github.com/kataras/iris"
)

func main() {

	iris.UseFunc(responseLogger) // global middleware, catch all routes

	iris.Get("/", func(c *iris.Context) {
		c.Write("Hello from %s", c.PathString())
	})

	iris.Get("/home", func(c *iris.Context) {
		c.Write("Hello from %s", c.PathString())
	})

	fmt.Println("Server is running at: 8080")
	iris.Listen(":8080")
}

func responseLogger(c *iris.Context) {
	c.Next() // process the request first, we don't want to have delays

	date := time.Now().Format("01/02 - 15:04:05")
	fmt.Printf("%s\n%s", date, c.Response.String())
}
