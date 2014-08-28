package main

import (
	"github.com/hoisie/web"
)

func welcome(ctx *web.Context, val string) {
	for k, v := range ctx.Params {
		println(k, v)
	}
}

func main() {
	web.Get("/(.*)", welcome)
	web.Run("0.0.0.0:9999")
}
