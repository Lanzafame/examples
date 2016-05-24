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
	w := ws.New()
	chat(w)
	iris.Static("/js", "./static/js", 1)

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Render("client.html", clientPage{"Client Page", ctx.HostString()})
	})

	iris.Get("/ws", func(ctx *iris.Context) {
		if err := w.Upgrade(ctx); err != nil {
			iris.Logger().Panic(err)
		}
	})

	fmt.Println("Server is listening at: 8080")
	iris.Listen(":8080")
}

func chat(w ws.Server) {

	w.OnConnection(func(c ws.Connection) {
		c.On("chat", func(message string) {
			c.To(ws.Broadcast).Emit("chat", "Native websocket message from: "+c.ID()+"-> "+message) // to all except this connection //worked
			c.Emit("chat", "to my self: "+message)
		})
	})

}
