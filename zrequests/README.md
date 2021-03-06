zrequests
=======

Simple and sane HTTP request library for Go language.


**Table of Contents**

- [Why zrequests?](#user-content-why-zrequests)
- [How do I install it?](#user-content-how-do-i-install-it)
- [What can I do with it?](#user-content-what-can-i-do-with-it)
  - [Making requests with different methods](#user-content-making-requests-with-different-methods)
  - [GET](#user-content-get)
  - [POST](#user-content-post)
    - [Sending payloads in the Body](#user-content-sending-payloads-in-the-body)
  - [Specifiying request headers](#user-content-specifiying-request-headers)
  - [Setting timeouts](#user-content-setting-timeouts)
 - [Using the Response and Error](#user-content-using-the-response-and-error)
 - [Receiving JSON](#user-content-receiving-json)
 - [Sending/Receiving Compressed Payloads](#user-content-sendingreceiving-compressed-payloads)
    - [Using gzip compression:](#user-content-using-gzip-compression)
    - [Using deflate compression:](#user-content-using-deflate-compression)
    - [Using compressed responses:](#user-content-using-compressed-responses)
 - [Proxy](#proxy)
 - [TODO:](#user-content-todo)


Why zrequests?
==========

Go has very nice native libraries that allows you to do lots of cool things. But sometimes those libraries are too low level, which means that to do a simple thing, like an HTTP Request, it takes some time. And if you want to do something as simple as adding a timeout to a request, you will end up writing several lines of code.

This is why we think zrequests is useful. Because you can do all your HTTP requests in a very simple and comprehensive way, while enabling you to do more advanced stuff by giving you access to the native API.

How do I install it?
====================

```bash
go get github.com/jack-zh/ztodo/zrequests
```

What can I do with it?
======================

## Making requests with different methods

#### GET
```go
res, err := zrequests.Request{ Uri: "http://www.google.com" }.Do()
```

zrequests default method is GET.

You can also set value to GET method easily

```go
type Item struct {
        Limit int
        Skip int
        Fields string
}

item := Item {
        Limit: 3,
        Skip: 5,
        Fields: "Value",
}

res, err := zrequests.Request{
        Uri: "http://localhost:3000/",
        QueryString: item,
}.Do()
```
The sample above will send `http://localhost:3000/?limit=3&skip=5&fields=Value`


QueryString also support url.Values

```go
item := url.Values{}
item.Set("Limit", 3)
item.Add("Field", "somefield")
item.Add("Field", "someotherfield")

res, err := zrequests.Request{
        Uri: "http://localhost:3000/",
        QueryString: item,
}.Do()
```

The sample above will send `http://localhost:3000/?limit=3&field=somefield&field=someotherfield`



#### POST

```go
res, err := zrequests.Request{ Method: "POST", Uri: "http://www.google.com" }.Do()
```

## Sending payloads in the Body

You can send ```string```, ```Reader``` or ```interface{}``` in the body. The first two will be sent as text. The last one will be marshalled to JSON, if possible.

```go
type Item struct {
    Id int
    Name string
}

item := Item{ Id: 1111, Name: "foobar" }

res, err := zrequests.Request{
    Method: "POST",
    Uri: "http://www.google.com",
    Body: item,
}.Do()
```

## Specifiying request headers

We think that most of the times the request headers that you use are: ```Host```, ```Content-Type```, ```Accept``` and ```User-Agent```. This is why we decided to make it very easy to set these headers.

```go
res, err := zrequests.Request{
    Uri: "http://www.google.com",
    Host: "foobar.com",
    Accept: "application/json",
    ContentType: "application/json",
    UserAgent: "zrequests",
}.Do()
```

But sometimes you need to set other headers. You can still do it.

```go
req := Request{ Uri: "http://www.google.com" }

req.AddHeader("X-Custom", "somevalue")

req.Do()
```

## Setting timeouts

zrequests supports 2 kind of timeouts. A general connection timeout and a request specific one. By default the connection timeout is of 1 second. There is no default for request timeout, which means it will wait forever.

You can change the connection timeout doing:

```go
zrequests.SetConnectTimeout(100 * time.Millisecond)
```

And specify the request timeout doing:

```go
res, err := zrequests.Request{
    Uri: "http://www.google.com",
    Timeout: 500 * time.Millisecond,
}.Do()
```

## Using the Response and Error

zrequests will always return 2 values: a ```Response``` and an ```Error```.
If ```Error``` is not ```nil``` it means that an error happened while doing the request and you shouldn't use the ```Response``` in any way.
You can check what happened by getting the error message:

```go
fmt.Println(err.Error())
```
And to make it easy to know if it was a timeout error, you can ask the error or return it:

```go
if serr, ok := err.(*zrequests.Error); ok {
    if serr.Timeout() {
        ...
    }
}
return err
```

If you don't get an error, you can safely use the ```Response```.

```go
res.StatusCode //return the status code of the response
res.Body // gives you access to the body
res.Body.ToString() // will return the body as a string
res.Header.Get("Content-Type") // gives you access to all the response headers
```
Remember that you should **always** close `res.Body` if it's not `nil`

## Receiving JSON

zrequests will help you to receive and unmarshal JSON.

```go
type Item struct {
    Id int
    Name string
}

var item Item

res.Body.FromJsonTo(&item)
```

## Sending/Receiving Compressed Payloads
zrequests supports gzip, deflate and zlib compression of requests' body and transparent decompression of responses provided they have a correct `Content-Encoding` header.

#####Using gzip compression:
```go
res, err := zrequests.Request{
    Method: "POST",
    Uri: "http://www.google.com",
    Body: item,
    Compression: zrequests.Gzip(),
}.Do()
```
#####Using deflate compression:
```go
res, err := zrequests.Request{
    Method: "POST",
    Uri: "http://www.google.com",
    Body: item,
    Compression: zrequests.Deflate(),
}.Do()
```
#####Using zlib compression:
```go
res, err := zrequests.Request{
    Method: "POST",
    Uri: "http://www.google.com",
    Body: item,
    Compression: zrequests.Zlib(),
}.Do()
```
#####Using compressed responses:
If servers replies a correct and matching `Content-Encoding` header (gzip requires `Content-Encoding: gzip` and deflate `Content-Encoding: deflate`) zrequests transparently decompresses the response so the previous example should always work:
```go
type Item struct {
    Id int
    Name string
}
res, err := zrequests.Request{
    Method: "POST",
    Uri: "http://www.google.com",
    Body: item,
    Compression: zrequests.Gzip(),
}.Do()
var item Item
res.Body.FromJsonTo(&item)
```
If no `Content-Encoding` header is replied by the server zrequests will return the crude response.

## Proxy
If you need to use a proxy for your requests zrequests supports the standard `http_proxy` env variable as well as manually setting the proxy for each request

```go
res, err := zrequests.Request{
    Method: "GET",
    Proxy: "http://myproxy:myproxyport",
    Uri: "http://www.google.com",
}.Do()
```

### Proxy basic auth is also supported

```go
res, err := zrequests.Request{
    Method: "GET",
    Proxy: "http://user:pass@myproxy:myproxyport",
    Uri: "http://www.google.com",
}.Do()
```
