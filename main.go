package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {
	loadEnvironment()

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"MapboxToken": os.Getenv("MAPBOX_TOKEN"),
		})
	})

	startServer(app)
}

func startServer(app *fiber.App) {
	port := os.Getenv("PORT")
	log.Fatal(app.Listen(port))
}

func loadEnvironment() {
	if os.Getenv("DEV") != "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}
