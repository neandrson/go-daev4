package calculator

import (
	"github.com/gofiber/fiber/v2"
	"github.com/neandrson/go-daev4/calculator/routes"
	// "github.com/google/uuid"
)

type Calculator struct {
	ServerId string `json:"server_id"`
}

func (o *Calculator) Start() {
	// o.ServerId = uuid.New().String()
	serverCalculator := fiber.New()

	routes.SetupRoutesCalculator(serverCalculator)
	serverCalculator.Listen(":8085")
}
