package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"os/exec"
	"playground/GoConvert/convert"
)

func setupRoutes(app *fiber.App) {

	app.Post("/api/v1/convert", convert.Convert)

}

func init() {
	go startLibre()
}

func main() {
	app := fiber.New()

	setupRoutes(app)
	err := app.Listen(4000)
	if err != nil {
		fmt.Println("Err listening on port:", err)
	}

}

func startLibre() {
	cmd := exec.Command("libreoffice", "--headless")
	if err := cmd.Run(); err != nil {
		fmt.Printf("error with libre: %v \n", err)
		fmt.Println("")
	}
}
