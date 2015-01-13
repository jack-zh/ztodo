package main

import (
	"fmt"
	"github.com/jack-zh/ztodo/zRequests"
)

func main() {
	fmt.Println("Hello")
	type Item struct {
		UserName string
		Password string
	}

	item := Item{UserName: "jack", Password: "123456"}
	res, err := zRequests.Request{
		Method:      "POST",
		Uri:         "http://localhost:9999/login",
		QueryString: item,
	}.Do()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.Body.ToString())
	}
}
