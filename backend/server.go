package main

import (
	"azorimor/azo-image-scraper/database"
	"embed"

	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Embed the frontend application
//
//go:embed ui/*
var frontendApplication embed.FS

func main() {
	dbSetup()

	app := fiber.New()
	app.Use(logger.New())

	app.Use("/ui", filesystem.New(filesystem.Config{
		Root:       http.FS(frontendApplication),
		PathPrefix: "ui",
		Browse:     true,
	}))

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
