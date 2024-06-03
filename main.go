package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/ragazoni/debt/internal/service"
)

func main() {

	app := fiber.New()

	v1 := app.Group("/v1")

	service.SetRouter(v1)

	app.Listen(":3000")

}
