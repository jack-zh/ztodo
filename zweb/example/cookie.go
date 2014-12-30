package main

import (
    "fmt"
    "github.com/jack-zh/ztodo/zweb"
    "html"
)

var cookieName = "cookie"

var notice = `
<div>%v</div>
`
var form = `
<form method="POST" action="update">
  <div class="field">
    <label for="cookie"> Set a cookie: </label>
    <input id="cookie" name="cookie"> </input>
  </div>

  <input type="submit" value="Submit"></input>
  <input type="submit" name="submit" value="Delete"></input>
</form>
`

func index(ctx *zweb.Context) string {
    cookie, _ := ctx.Request.Cookie(cookieName)
    var top string
    if cookie == nil {
        top = fmt.Sprintf(notice, "The cookie has not been set")
    } else {
        var val = html.EscapeString(cookie.Value)
        top = fmt.Sprintf(notice, "The value of the cookie is '"+val+"'.")
    }
    return top + form
}

func update(ctx *zweb.Context) {
    if ctx.Params["submit"] == "Delete" {
        ctx.SetCookie(zweb.NewCookie(cookieName, "", -1))
    } else {
        ctx.SetCookie(zweb.NewCookie(cookieName, ctx.Params["cookie"], 0))
    }
    ctx.Redirect(301, "/")
}

func main() {
    zweb.Get("/", index)
    zweb.Post("/update", update)
    zweb.Run("0.0.0.0:9999")
}
