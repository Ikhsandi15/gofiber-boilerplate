package main

import (
	"flag"
	"gofiber_boilerplate/config"
	"gofiber_boilerplate/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var m bool

	config.Connection()

	flag.BoolVar(
		&m,
		"m",
		false,
		`This flag is used for first migration`,
	)

	flag.Parse()

	if m {
		config.AutoMigrate(config.DB)
		return
	}

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":3000")
}
