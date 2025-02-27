package main

import (
	calculator "github.com/neandrson/go-daev4/calculator/server"
)

func main() {
	calculator := &calculator.Calculator{}

	calculator.Start()
}
