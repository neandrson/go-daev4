package main

import (
	"github.com/neandrson/go-daev4/orchestrator/app"
)

func main() {
	orchestrator := &app.Orchestrator{}
	orchestrator.Run()
}
