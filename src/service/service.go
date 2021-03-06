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

//Adds handler for command
func (s Service) AddHandler(command string, handler func(server.ClientListener, []string)) {
	s.RequestHandlers[command] = handler
}

//Handles command from parsed client's message
func (s Service) HandleCommand(c server.ClientListener, message string) {
	if s.HandleCommandMethod == nil {
		s.defaultCommandHandler(c, message)
		return
	}
	s.HandleCommandMethod(s, c, message)
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
	handler, ok := s.RequestHandlers[cmd]

	if !ok {
		c.SendError("invalid command " + cmd)
		return
	}
	handler(c, args)
}
