package helloworld

import (
	"fmt"
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

// NewService returns a helloworld service instance
func NewService() Service {
	return Service{}
}

////////////////////////////////////////////////////////
// Service tasks
////////////////////////////////////////////////////////

func replyHello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}
