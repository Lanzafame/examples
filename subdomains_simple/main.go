package main

import (
	"github.com/kataras/iris"
)

func main() {
	api := iris.New()

	// first the subdomains.
	admin := api.Party("admin.yourhost.com")
	{
		//this will only success on admin.yourhost.com/hey
		admin.Get("/hey", func(c *iris.Context) {
			c.Write("HEY FROM admin.yourhost.com")
		})
		//this will only success on admin.yourhost.com/hey2
		admin.Get("/hey2", func(c *iris.Context) {
			c.Write("HEY SECOND FROM admin.yourhost.com")
		})
	}

	api.Get("/hey", func(c *iris.Context) {
		c.Write("HEY FROM no-subdomain hey")
	})

	println("Server is listening at: 80")
	api.Listen(":80")
}
