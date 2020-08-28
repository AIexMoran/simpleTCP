package handler

type ClientListenerInterface interface {
	SendMessage(string)
	SendError(string)
	ReadMessage() (string, bool)
}
