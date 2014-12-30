package main

import (
    "github.com/jack-zh/ztodo/zweb"
    "net/http"
    "strconv"
    "time"
)

func hello(ctx *zweb.Context, num string) {
    flusher, _ := ctx.ResponseWriter.(http.Flusher)
    flusher.Flush()
    n, _ := strconv.ParseInt(num, 10, 64)
    for i := int64(0); i < n; i++ {
        ctx.WriteString("<br>hello world</br>")
        flusher.Flush()
        time.Sleep(1e9)
    }
}

func main() {
    zweb.Get("/([0-9]+)", hello)
    zweb.Run("0.0.0.0:9999")
}
