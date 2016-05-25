package main

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/ws"
)

type clientPage struct {
	Title string
	Host  string
}

func main() {
	api := iris.New()

	api.Static("/js", "./static/js", 1)

	api.Get("/", func(ctx *iris.Context) {
		ctx.Render("client.html", clientPage{"Client Page", ctx.HostString()})
	})

	// important staff

	w := ws.New(api, "/my_endpoint")
	// for default 'iris.' station use that: w := ws.New(iris.DefaultIris, "/my_endpoint")

	w.OnConnection(func(c ws.Connection) {
		c.On("chat", func(message string) {
			c.To(ws.Broadcast).Emit("chat", "Message from: "+c.ID()+"-> "+message) // to all except this connection
			c.Emit("chat", "Message from myself: "+message)
		})
	})

	//

	fmt.Println("Server is listening at: 8080")
	api.Listen(":8080")
}
