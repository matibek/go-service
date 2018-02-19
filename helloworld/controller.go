package helloworld

import (
	"net/http"

	"github.com/matibek/service-scaffolding-go/core"
)

// ReplyHello will respond hello message
func ReplyHello(c *core.Context) {
	name := c.DefaultQuery("name", "World")
	result := replyHello(name)
	c.String(http.StatusOK, result)
}
