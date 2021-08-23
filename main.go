package main

import (
	"fmt"
	"os"
	"log"
	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/RohitKuwar/go-fiber-api/routes"
)

func setupRoutes(app *fiber.App) {
	// give response when at /
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success":  true,
			"message": "You are at the endpoint 😉",
		})
	})

	// api group
	api := app.Group("/api")

	// give response when at /api
	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the api endpoint 😉",
		})
	})

	// send todos route group to TodoRoutes of routes package
	routes.TodoRoute(api.Group("/todos"))
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// setup routes
	setupRoutes(app)

	// Listen on server 8000 and catch error if any
	// port := os.Getenv("PORT")
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  port := os.Getenv("PORT")
	app.Listen(":"+port)
	fmt.Println("port is:", port)

}