package helloworld

import "github.com/matibek/go-service/core"

// RegisterRoute adds a routing to the driver
func RegisterRoute(driver *core.Engine) {
	driver.GET("/helloworld", ReplyHello)
}
