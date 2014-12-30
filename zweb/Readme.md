# zweb.go

zweb.go is the simplest way to write web applications in the Go programming language. It's ideal for writing simple, performant backend web services for ztodo, and it frork from web.go. 

## Overview

zweb.go should be familiar to people who've developed websites with higher-level web frameworks like sinatra or tornado. It is designed to be a lightweight web framework that doesn't impose any scaffolding on the user. Some features include:

* Routing to url handlers based on regular expressions
* Secure cookies
* Support for fastcgi and scgi
* Web applications are compiled to native code. This means very fast execution and page render speed
* Efficiently serving static files

## Installation


To install zweb.go, simply run:

    go get github.com/jack-zh/ztodo/zweb

## Example
```go
package main
    
import (
    "github.com/jack-zh/ztodo/zweb"
)
    
func hello(val string) string { return "hello " + val } 
    
func main() {
    zweb.Get("/(.*)", hello)
    zweb.Run("0.0.0.0:9999")
}
```

To run the application, put the code in a file called hello.go and run:

    go run hello.go
    
You can point your browser to http://localhost:9999/world . 

### Getting parameters

Route handlers may contain a pointer to web.Context as their first parameter. This variable serves many purposes -- it contains information about the request, and it provides methods to control the http connection. For instance, to iterate over the web parameters, either from the URL of a GET request, or the form data of a POST request, you can access `ctx.Params`, which is a `map[string]string`:

```go
package main

import (
    "github.com/jack-zh/ztodo/zweb"
)
    
func hello(ctx *zweb.Context, val string) { 
    for k,v := range ctx.Params {
		println(k, v)
	}
}
    
func main() {
    zweb.Get("/(.*)", hello)
    zweb.Run("0.0.0.0:9999")
}
```

In this example, if you visit `http://localhost:9999/?a=1&b=2`, you'll see the following printed out in the terminal:

    a 1
    b 2

## Documentation

the code.

## About

web.go was written by [Jack.z](http://link-pub.cn). 

## License:
    
    MIT


