package main

import (
	"azorimor/azo-image-scraper/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	dbSetup()

	app := fiber.New()
	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":8111")
}

func dbSetup() {
	err := database.CreateDBConnection()
	if err != nil {
		panic(err)
	}
	err = database.AutoMigrateDB()
	if err != nil {
		panic(err)
	}
}
