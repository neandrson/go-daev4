package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/neandrson/go-daev4/orchestrator/routes"
)

type Orchestrator struct{}

func (o *Orchestrator) Run() {
	server := fiber.New()
	server.Use(cors.New())

	routes.Setup(server)
	server.Listen(":8081")
}
