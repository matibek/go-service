package helloworld

import "github.com/matibek/service-scaffolding-go/core"

// Service for helloworld
type Service struct{}

// RegisterRoute adds a routing to the driver
func (Service) RegisterRoute(serviceDriver *core.Engine) {
	serviceDriver.GET("/helloworld", ReplyHello)
}

// NewService returns a helloworld service instance
func NewService() Service {
	return Service{}
}
