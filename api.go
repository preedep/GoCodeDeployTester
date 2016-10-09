package main

import (
	"fmt"

	"github.com/kataras/iris"
)

func main() {
	iris.Get("/hi", func(ctx *iris.Context) {
		ctx.Write("Hello world %s", "iris")

	})
	fmt.Println("Start Service....")
	iris.Listen(":8080")
}
