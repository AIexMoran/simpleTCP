package main

import (
	"github.com/aIexmoran/simpletcp/src/handler"
	"github.com/aIexmoran/simpletcp/src/server"
	"github.com/aIexmoran/simpletcp/src/service"
	"os"
)

var port string

func init() {
	port = os.Getenv("SERVER_PORT")

	if port == "" {
		panic("Set port!")
	}
}

func main() {
	serv := service.CreateService(handler.TestProtocol, handler.TestHandler)

	serv.AddHandler("time", handler.TimeHandler)
	serv.AddHandler("print", handler.PrintHandler)
	serv.AddHandler("hello", handler.HelloHandler)
	server.ListenAndHandle(serv, "tcp", ":"+port)
}
