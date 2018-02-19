package helloworld

import (
	"fmt"
)

func replyHello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}
