package main

import (
    "github.com/jack-zh/ztodo/zweb"
    "log"
    "os"
)

func hello(val string) string { return "hello " + val }

func main() {
    f, err := os.Create("server.log")
    if err != nil {
        println(err.Error())
        return
    }
    logger := log.New(f, "", log.Ldate|log.Ltime)
    zweb.Get("/(.*)", hello)
    zweb.SetLogger(logger)
    zweb.Run("0.0.0.0:9999")
}
