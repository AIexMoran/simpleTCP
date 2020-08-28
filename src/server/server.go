package server

import (
	"bufio"
	"github.com/aIexmoran/simpletcp/src/client"

	"net"
)

type ClientListener interface {
	SendMessage(string)
	SendError(string)
	ReadMessage() (string, bool)
}

type RequestHandler interface {
	HandleCommand(ClientListener, string)
}

func handleClient(conn net.Conn, handler RequestHandler) {
	defer conn.Close()

	cli := client.Client{
		Writer: bufio.NewWriter(conn),
		Reader: bufio.NewReader(conn),
	}
	for {
		message, ok := cli.ReadMessage()
		if !ok {
			return
		}
		handler.HandleCommand(cli, message)
	}
}

//Listens on the protocol, port
func ListenAndHandle(handler RequestHandler, protocol, port string) error {
	ln, err := net.Listen(protocol, port)

	if err != nil {
		return err
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn, handler)
	}
}
