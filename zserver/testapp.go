package main

import (
	"fmt"
	"github.com/jack-zh/ztodo/zrequests"
)

type Item struct {
	UserName string
	Password string
}

func signup() string {
	return "signup-->"
}

func login() string {
	return ""
}

func getuser(usertokenstr string) string {
	return "getuser--> usertokenstr" + usertokenstr
}

func pullall() {
	// item := Item{UserName: "jack", Password: "123456"}
	res, err := zrequests.Request{
		Method: "GET",
		Uri:    "http://localhost:9999/pullall/aaabbbcccddd",
		// QueryString: item,
	}.Do()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.Body.ToString())
	}
}

func pullone() {
	// item := Item{UserName: "jack", Password: "123456"}
	res, err := zrequests.Request{
		Method: "GET",
		Uri:    "http://localhost:9999/pullall/aaabbbcccddd",
		// QueryString: item,
	}.Do()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.Body.ToString())
	}
}

func pushall(usertokenstr string) string {
	return "pushall==> usertokenstr:" + usertokenstr
}

func pushone(usertokenstr string) string {
	return "pushone==> usertokenstr:" + usertokenstr
}

func pnf(url string) string {
	return "{'error': 404}"
}

func main() {
	fmt.Println("Hello")

}
