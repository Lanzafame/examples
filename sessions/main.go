package main

import (
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"

	_ "github.com/kataras/iris/sessions/providers/memory" // add a store it is auto-registers itself
)

var sess *sessions.Manager

func init() {
	sess = sessions.New("memory", "irissessionid", time.Duration(time.Minute)*60)
}

func main() {

	iris.Get("/set", func(c *iris.Context) {
		//get the session for this context
		session := sess.Start(c)

		//set session values
		session.Set("name", "kataras")

		//test if setted here
		c.Write("All ok session setted to: %s", session.Get("name"))
	})

	iris.Get("/get", func(c *iris.Context) {
		//get the session for this context
		session := sess.Start(c)

		var name string

		//get the session value
		if v := session.Get("name"); v != nil {
			name = v.(string)
		}
		// OR just name = session.GetString("name")

		c.Write("The name on the /set was: %s", name)
	})

	iris.Get("/clear", func(c *iris.Context) {
		//get the session for this context
		session := sess.Start(c)

		session.Delete("name")

	})

	println("Server is listening at :8080")
	iris.Listen("8080")

}
