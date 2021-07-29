package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"playground/GoConvert/convert"
)

func setupRoutes(app *fiber.App) {

	app.Post("/api/v1/convert", convert.Convert)

}

func main() {
	app := fiber.New()

	setupRoutes(app)
	err := app.Listen(4000)
	if err != nil {
		fmt.Println("Err listening on port:", err)
	}

}
