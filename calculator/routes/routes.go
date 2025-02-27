package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/neandrson/go-daev4/calculator/handlers"
)

func SetupRoutesCalculator(app *fiber.App) {
	handlers := &handlers.CalculatorHendler{}

	app.Get("/start", handlers.Start)
}
