package main

import (
	"github.com/kataras/iris"
)

type MyData struct {
	Sysname   string // this will be the same for all requests
	Version   int    // this will be the same for all requests
	UserAgent string // this will be different for each request
}

type MyHandler struct {
	data MyData
}

func (m *MyHandler) Serve(ctx *iris.Context) {
	data := &m.data
	data.UserAgent = ctx.RequestHeader("User-Agent")
	ctx.Write("User agent: %s", data.UserAgent)
}

func main() {
	/*
	 Create the data with the predefined fields that will no change at every request
	*/
	myData := MyData{
		Sysname: "Redhat",
		Version: 1,
	}
	/*
		use a new MyHandler for each route,
		keep the same original myData with the fields no need to change on every request
	*/
	iris.Handle("GET", "/", &MyHandler{myData})
	iris.Handle("GET", "/about", &MyHandler{myData})

	iris.Listen(":8080")
}
