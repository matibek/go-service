package helloworld

import (
	"net/http"

	"github.com/matibek/go-service/core"
)

// ReplyHello will respond hello message
func ReplyHello(c *core.Context) {
	name := c.DefaultQuery("name", "World")
	result, err := replyHello(name)
	if err != nil {
		err := core.Errors.NewWithMessage(err, "This sample error to client")
		c.Error(err)
		return
	}
	c.String(http.StatusOK, result)
}
