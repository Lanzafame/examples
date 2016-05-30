package main

import (
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/config"
	"github.com/kataras/iris/middleware/basicauth"
)

func main() {
	authConfig := config.BasicAuth{
		Users:      map[string]string{"myusername": "mypassword", "mySecondusername": "mySecondpassword"},
		Realm:      "Authorization Required", // if you don't set it it's "Authorization Required"
		ContextKey: "user",                   // if you don't set it it's "auth"
		Expires:    time.Duration(30) * time.Minute,
	}

	authentication := basicauth.New(authConfig)

	// to global iris.UseFunc(authentication)
	// to routes
	/*
		iris.Get("/mysecret", authentication, func(ctx *iris.Context) {
			username := ctx.GetString("user") //  the Contextkey from the authConfig
			ctx.Write("Hello authenticated user: %s ", username)
		})
	*/

	// to party

	needAuth := iris.Party("/secret", authentication)
	{
		needAuth.Get("/", func(ctx *iris.Context) {
			username := ctx.GetString("user") //  the Contextkey from the authConfig
			ctx.Write("Hello authenticated user: %s from localhost:8080/secret ", username)
		})

		needAuth.Get("/profile", func(ctx *iris.Context) {
			username := ctx.GetString("user") //  the Contextkey from the authConfig
			ctx.Write("Hello authenticated user: %s from localhost:8080/secret/profile ", username)
		})

		needAuth.Get("/settings", func(ctx *iris.Context) {
			username := ctx.GetString("user") //  the Contextkey from the authConfig
			ctx.Write("Hello authenticated user: %s from localhost:8080/secret/settings ", username)
		})
	}

	iris.Listen(":8080")
}
