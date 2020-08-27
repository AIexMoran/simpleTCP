package handler

import (
	"errors"
	"fmt"
	"github.com/aIexmoran/simpletcp/src/server"
	"github.com/aIexmoran/simpletcp/src/service"
	"strings"
)

func TestProtocol(message string) (string, []string, error) {
	request := strings.SplitN(strings.TrimSpace(message), ":", 2)

	if len(request) != 2 {
		return "", []string{}, errors.New("usage: cmd: arg1 arg2 ... argN")
	}
	return request[0], strings.Split(strings.TrimSpace(request[1]), " "), nil
}

func TestHandler(s service.Service, c server.ClientListener, message string) {
	fmt.Println("This is custom handler")
	if len(message) > 30 {
		c.SendError("Too many bytes")
		return
	}
	cmd, args, err := s.ProtocolMethod(message)

	if err != nil {
		c.SendError(err.Error())
		return
	}
	val, ok := s.RequestHandlers[cmd]

	if !ok {
		c.SendError("invalid command " + cmd)
		return
	}
	val(c, args)
}
