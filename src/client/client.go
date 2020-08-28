package client

import (
	"bufio"
	"log"
)

type Client struct {
	Writer *bufio.Writer
	Reader *bufio.Reader
}

//Sends message to client
func (c Client) SendMessage(message string) {
	c.Writer.WriteString(message + "\n")
	c.Writer.Flush()
}

//Sends error to client
func (c Client) SendError(message string) {
	c.SendMessage("Error:" + message)
}

//Reads message from client
func (c Client) ReadMessage() (string, bool) {
	message, err := c.Reader.ReadString('\n')

	if err != nil {
		log.Println(err)
		return "", false
	}
	return message, true
}
