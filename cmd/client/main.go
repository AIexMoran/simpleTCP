package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var port string
var protocol = "tcp"

func init() {
	port = os.Getenv("PORT")

	if port == "" {
		panic("Set PORT!")
	}
	port = ":" + port
}

func handleServerAnswers(reader *bufio.Reader) {
	var (
		answer string
		err    error
	)

	for {
		answer, err = reader.ReadString('\n')

		if err != nil {
			panic("Cannot get answer")
		}
		fmt.Println("\nSERVER: " + answer)
	}
}

func main() {
	fmt.Println("Connecting to port", port[1:], "protocol", protocol)
	conn, err := net.Dial(protocol, port)
	writer := bufio.NewWriter(conn)
	serverReader := bufio.NewReader(conn)
	reader := bufio.NewReader(os.Stdin)

	if err != nil {
		fmt.Println(err)
		panic("Could't connect to server")
	}
	go handleServerAnswers(serverReader)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			panic("Cannot read string")
		}
		_, err = writer.WriteString(message)
		writer.Flush()
		if err != nil {
			panic("Cannot send message")
		}
	}
}
