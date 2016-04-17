package main

import (
	"github.com/kataras/iris"
)

func main() {
	api := iris.New()
	//api := iris.Custom(iris.StationOptions{Cache: false})
	//ok it's working both cached and normal router, optimization done before listen

	//the subdomains are working like parties, the only difference is that you CANNOT HAVE party of party with both of them subdomains
	//YOU CANNOT DO THAT AND YOU MUSTN'T DO IT DIRECTLY: api.Party("admin.yourhost.com").Party("other.admin.yuorhost.com")
	//Do that: api.Party("other.admin.yourhost.com") .... and different/new party with api.Party("admin.yourhost.com")
	admin := api.Party("admin.yourhost.com")
	{
		//this will only success on admin.yourhost.com/hey
		admin.Get("/hey", func(c *iris.Context) {
			c.Write("HEY FROM admin.omicronware.com")
		})
		//this will only success on admin.yourhost.com/hey2
		admin.Get("/hey2", func(c *iris.Context) {
			c.Write("HEY SECOND FROM admin.omicronware.com")
		})
	}

	// this will serve on yourhost.com/hey and not on admin.yourhost.com/hey
	api.Get("/hey", func(c *iris.Context) {
		c.Write("HEY FROM no-subdomain hey")
	})

	api.Listen(":80")
}
