package main

import (
	"github.com/gofiber/fiber/v2"
	"os"
)

var data []interface{}

func setupRouter() *fiber.App {
	app := &fiber.App{}

	app.Get("serve", func(ctx *fiber.Ctx) error {
		return ctx.JSON("data")
	})

	return app
}

func setupData(file os.File) {
}

func check(error error) {
	if error != nil {
		panic(error)
	}
}

func main() {
	app := setupRouter()
	_ = app.Listen(":5000")

}
