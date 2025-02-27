package main

import (
	"github.com/evsedov/GoCalculator/web/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
)

func main() {

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})
	app.Use(cors.New(cors.ConfigDefault))
	app.Static("/", "./static")
	routes.SetupRoutes(app)
	app.Listen(":8080")
}
