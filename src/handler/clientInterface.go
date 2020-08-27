package handler

type ClientListenerInterface interface {
	SendMessage(string)
	ReadMessage() (string, bool)
}
