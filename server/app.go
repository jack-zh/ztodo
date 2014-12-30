package main
    
import (
    "github.com/jack-zh/ztodo/zweb"
)
    
func hello(val string) string { return "hello " + val } 
    
func main() {
    zweb.Get("/(.*)", hello)
    zweb.Run("0.0.0.0:9999")
}
