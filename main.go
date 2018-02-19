package main

import (
	"github.com/matibek/service-scaffolding-go/core"
	"github.com/matibek/service-scaffolding-go/helloworld"
)

// Application entry point
func main() {
	core.Logger.Info("Starting server...")
	helloworldService := helloworld.NewService()
	server := core.NewServer(helloworldService)
	server.Start()
}
