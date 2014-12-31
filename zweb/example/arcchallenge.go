package main

import (
	"fmt"
	"github.com/jack-zh/ztodo/zweb"
	"math/rand"
	"time"
)

var form = `<form action="say" method="POST"><input name="said"><input type="submit"></form>`

var users = map[string]string{}

func main() {
	rand.Seed(time.Now().UnixNano())
	zweb.Config.CookieSecret = "7C19QRmwf3mHZ9CPAaPQ0hsWeufKd"
	zweb.Get("/", func(ctx *zweb.Context) string {
		ctx.Redirect(302, "/said")
		return ""
	})
	zweb.Get("/said", func() string { return form })
	zweb.Post("/say", func(ctx *zweb.Context) string {
		uid := fmt.Sprintf("%d\n", rand.Int63())
		ctx.SetSecureCookie("user", uid, 3600)
		users[uid] = ctx.Params["said"]
		return `<a href="/final">Click Here</a>`
	})
	zweb.Get("/final", func(ctx *zweb.Context) string {
		uid, _ := ctx.GetSecureCookie("user")
		return "You said " + users[uid]
	})
	zweb.Run("0.0.0.0:9999")
}
