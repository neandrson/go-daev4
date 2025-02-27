package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/neandrson/go-daev4/web/handlers"
)

func SetupRoutes(app *fiber.App) {
	mainHandler := &handlers.MainHandler{}

	app.Get("/", mainHandler.GetHome)
	app.Get("/login", mainHandler.GetLogin)
	app.Get("/register", mainHandler.GetRegister)
}
