package main

import (
	"fmt"
	"github.com/jack-zh/ztodo/zrequests"
)

type Item struct {
	UserName string
	Password string
}

type BackCode struct {
	Message string
	CodeNum int
}

func signup(username string, password string) BackCode {
	item := Item{UserName: username, Password: password}
	res, err := zrequests.Request{
		Method: "GET",
		Uri:    "http://localhost:9999/signup",
		QueryString: item,
	}.Do()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.Body.ToString())
	}
	return new BackCode{"success", 1}
}

func login(username string, password string) BackCode {
	item := Item{UserName: username, Password: password}
	res, err := zrequests.Request{
		Method: "GET",
		Uri:    "http://localhost:9999/login",
		QueryString: item,
	}.Do()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.Body.ToString())
	}
	return new BackCode{"success", 1}
}

func getuser(usertokenstr string) string {
	return "getuser--> usertokenstr" + usertokenstr
}

func pullall(usertokenstr string) {
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

func pullone(usertokenstr string, tasktokenstr string) {
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

func pushone(usertokenstr string, tasktokenstr string) string {
	return "pushone==> usertokenstr:" + usertokenstr
}

func pnf(url string) string {
	return "{'error': 404}"
}

func main() {
	fmt.Println("Hello")

}
