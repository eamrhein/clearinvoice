package main

import (
	"clearinvoice/db"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet/v2"
)

func main() {
	engine := jet.New("./views", ".jet")
	engine.Reload(true) // Enable template reloading in development mode
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	db.ConnectDatabase()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Clear Invoice",
			"Year":  time.Now().Year(),
		}, "layouts/main")
	})

	// Define routes...
	app.Listen(":3000")
}
