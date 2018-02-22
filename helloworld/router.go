package helloworld

import "github.com/matibek/service-scaffolding-go/core"

// RegisterRoute adds a routing to the driver
func RegisterRoute(driver *core.Engine) {
	driver.GET("/helloworld", ReplyHello)
}
