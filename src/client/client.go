package client

import (
	"bufio"
	"log"
)

type Client struct {
	Writer *bufio.Writer
	Reader *bufio.Reader
}

func (c Client) SendMessage(message string) {
	c.Writer.WriteString(message + "\n")
	c.Writer.Flush()
}

func (c Client) SendError(message string) {
	c.SendMessage("Error:" + message)
}

func (c Client) ReadMessage() (string, bool) {
	message, err := c.Reader.ReadString('\n')

	if err != nil {
		log.Println(err)
		return "", false
	}
	return message, true
}
