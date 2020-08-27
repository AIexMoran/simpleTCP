package service

import (
	"github.com/aIexmoran/simpletcp/src/server"
)

type Service struct {
	RequestHandlers     map[string]func(server.ClientListener, []string)
	ProtocolMethod      func(string) (string, []string, error)
	HandleCommandMethod func(Service, server.ClientListener, string)
}

func NewService(protocol func(string) (string, []string, error),
	handler func(Service, server.ClientListener, string)) *Service {
	return &Service{
		RequestHandlers:     make(map[string]func(server.ClientListener, []string)),
		ProtocolMethod:      protocol,
		HandleCommandMethod: handler,
	}
}

func (s Service) AddHandler(command string, handler func(server.ClientListener, []string)) {
	s.RequestHandlers[command] = handler
}

func (s Service) HandleCommand(c server.ClientListener, message string) {
	if s.HandleCommandMethod != nil {
		s.HandleCommandMethod(s, c, message)
		return
	}
	s.defaultCommandHandler(c, message)
}

func (s Service) defaultCommandHandler(c server.ClientListener, message string) {
	var args []string
	var err error
	cmd := message

	if s.ProtocolMethod != nil {
		cmd, args, err = s.ProtocolMethod(message)
		if err != nil {
			c.SendError(err.Error())
			return
		}
	}
	val, ok := s.RequestHandlers[cmd]

	if !ok {
		c.SendError("invalid command " + cmd)
		return
	}
	val(c, args)
}
