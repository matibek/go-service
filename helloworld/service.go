package helloworld

import (
	"errors"
	"fmt"

	"github.com/matibek/go-service/core"
)

// Service for helloworld
type Service struct{}

// Clean will do cleaning task before server exits
func (Service) Clean() (err error) {
	// TODO: put cleaning task here
	return
}

// Health will check the health of the service
func (Service) Health() (err error) {
	// TODO: check health (like connection)
	return
}

// RegisterRoute adds a routing to the driver
func (Service) RegisterRoute(driver *core.Engine) {
	RegisterRoute(driver)
}

// NewService returns a helloworld service instance
func NewService() Service {
	return Service{}
}

////////////////////////////////////////////////////////
// Service tasks
////////////////////////////////////////////////////////

func replyHello(name string) (string, error) {
	if name == "error" {
		return "", errors.New("This is inner error")
	}
	return fmt.Sprintf("Hello %s", name), nil
}
