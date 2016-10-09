package main

import (
	"fmt"

	"github.com/kataras/iris"
)

func main() {
	iris.Get("/hi", func(ctx *iris.Context) {
		ctx.SetCookieKV("name", "iris")
		ctx.Write("Hello world %s", ctx.GetCookie("name"))
	})
	fmt.Println("Start Service....")
	iris.Listen(":8080")
}
