package main

import (
	"math/rand"
	"time"

	"github.com/ironarachne/placenamegen"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		rand.Seed(time.Now().UnixNano())
		placeName := placenamegen.Generate()
		ctx.Write(placeName)
	})

	app.Run(iris.Addr(":7917"))
}
