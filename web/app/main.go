package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	"github.com/neandrson/go-daev4/web/routes"
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
