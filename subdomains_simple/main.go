package main

import (
	"github.com/kataras/iris"
)

func main() {
	api := iris.New()

	// first the subdomains.
	admin := api.Party("admin.yourhost.com")
	{
		// admin.yourhost.com
		admin.Get("/", func(c *iris.Context) {
			c.Write("HEY FROM admin.yourhost.com")
		})
		// admin.yourhost.com/hey
		admin.Get("/hey", func(c *iris.Context) {
			c.Write("HEY FROM admin.yourhost.com/hey")
		})
		// admin.yourhost.com/hey2
		admin.Get("/hey2", func(c *iris.Context) {
			c.Write("HEY SECOND FROM admin.yourhost.com/hey")
		})
	}

	// localhost/hey
	api.Get("/hey", func(c *iris.Context) {
		c.Write("HEY FROM no-subdomain hey")
	})

	api.Listen(":80")
}
