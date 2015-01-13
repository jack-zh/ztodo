package main

import (
	"fmt"
	"github.com/jack-zh/ztodo/zweb"
)

func signup(ctx *zweb.Context) string {
	return "signup-->"
}

func login(ctx *zweb.Context) string {
	for k, v := range ctx.Params {
		fmt.Println(k, v)
	}
	fmt.Println("...")

	// val := ctx.Request.MultipartForm.Value
	// for k1, v1 := range val {
	fmt.Println(ctx.Params)
	// }
	fmt.Println("...")
	return "login==>"
}

func getuser(usertokenstr string) string {
	return "getuser--> usertokenstr" + usertokenstr
}

func pullall(usertokenstr string) string {
	return "pullall --> usertokenstr:" + usertokenstr
}

func pullone(tasktoken string, usertokenstr string) string {
	return "pullone--> usertokenstr:" + usertokenstr + " | tasktoken:" + tasktoken
}

func pushall(ctx *zweb.Context, usertokenstr string) string {
	return "pushall==> usertokenstr:" + usertokenstr
}

func pushone(ctx *zweb.Context, usertokenstr string) string {
	return "pushone==> usertokenstr:" + usertokenstr
}

func pnf(url string) string {
	return "{'error': 404}"
}

func main() {
	zweb.Get("/pullall/(.*)", pullall)
	zweb.Get("/getuser/(.*)", getuser)
	zweb.Get("/pullone/(.*)/(.*)", pullone)

	zweb.Post("/signup", signup)
	zweb.Post("/login", login)

	zweb.Post("/pushall/(.*)", pushall)
	zweb.Post("/pushone/(.*)", pushone)

	zweb.Get("/(.*)", pnf)
	zweb.Post("/(.*)", pnf)

	zweb.Run("0.0.0.0:9999")
}
