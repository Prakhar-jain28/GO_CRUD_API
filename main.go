package main

import "gofr.dev/pkg/gofr"

func main() {
	app := gofr.New()

	app.GET("/home", func(ctx *gofr.Context) (interface{}, error) {
		return "Home Page", nil
	})

	app.Start()
	// app.Run("localhost:8080")
}