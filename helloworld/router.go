package helloworld

import "github.com/matibek/service-scaffolding-go/core"

// RegisterRoute adds a routing to the driver
func (Service) RegisterRoute(serviceDriver *core.Engine) {
	serviceDriver.GET("/helloworld", ReplyHello)
}
