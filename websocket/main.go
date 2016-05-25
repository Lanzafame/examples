package main

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/websocket"
)

type clientPage struct {
	Title string
	Host  string
}

/* All commented code is working*/

func main() {
	api := iris.New()

	api.Static("/js", "./static/js", 1)

	api.Get("/", func(ctx *iris.Context) {
		ctx.Render("client.html", clientPage{"Client Page", ctx.HostString()})
	})

	// important staff

	w := websocket.New(api, "/my_endpoint")
	// for default 'iris.' station use that: w := websocket.New(iris.DefaultIris, "/my_endpoint")
	var myChatRoom = "room1"
	w.OnConnection(func(c websocket.Connection) {

		/*c.On("join", func() {
			w.Join(myChatRoom, c) // care, no c.Join, but w.Join
		})

		c.On("leave", func() {
			w.Leave(myChatRoom, c)
		})*/

		w.Join(myChatRoom, c)

		c.On("chat", func(message string) {

			//c.To(websocket.Broadcast).Emit("chat", "Message from: "+c.ID()+"-> "+message) // to all except this connection
			//c.Emit("chat", "Message from myself: "+message)

			//send the message to the whole room, all connections are inside this room will receive this message
			c.To(myChatRoom).Emit("chat", "From: "+c.ID()+": "+message)
		})

		c.OnDisconnect(func() {
			fmt.Printf("\nConnection with ID: %s has been disconnected!", c.ID())
		})
	})

	//

	fmt.Println("Server is listening at: 8080")
	api.Listen(":8080")
}
