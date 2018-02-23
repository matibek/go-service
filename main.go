package main

import (
	"github.com/matibek/scaffold-go-service/core"
	"github.com/matibek/scaffold-go-service/helloworld"
)

// Application entry point
func main() {
	core.Logger.Info("Starting server...")
	helloworldService := helloworld.NewService()
	server := core.NewServer(helloworldService)
	server.Start()
}
