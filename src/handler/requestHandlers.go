package handler

import (
	"fmt"
	"github.com/aIexmoran/simpletcp/src/server"
	"strings"
	"time"
)

func PrintHandler(c server.ClientListener, args []string) {
	if len(args) < 2 {
		c.SendMessage("Usage: print: 'fmt' arg1 arg2 ... argn")
		return
	}
	fmtArgs := make([]interface{}, len(args)-1)
	for i, val := range args[1:] {
		fmtArgs[i] = val
	}
	c.SendMessage(fmt.Sprintf(args[0], fmtArgs...))
}

func TimeHandler(c server.ClientListener, args []string) {
	c.SendMessage(time.Now().String())
}

func HelloHandler(c server.ClientListener, args []string) {
	result := strings.Title(strings.Join(args, ", "))
	c.SendMessage("Hello! " + result)
}
