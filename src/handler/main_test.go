package handler

import (
	mock "github.com/aIexmoran/simpletcp/src/handler/mock"
	"github.com/aIexmoran/simpletcp/src/server"
	"github.com/golang/mock/gomock"
	"testing"
)

var tests = []struct {
	args []string
	want string
	handler func(server.ClientListener, []string)
}{
	{
		[]string{"1%s%s%s", "2", "3", "45"},
		"12345",
		PrintHandler,
	},
	{
		[]string{"1%s%s", "2", "34"},
		"1234",
		PrintHandler,
	},
	{
		[]string{"%s%s%s", "1", "2", "3"},
		"123",
		PrintHandler,
	},
	{
		[]string{"%s", "12"},
		"12",
		PrintHandler,
	},
	{
		[]string{"", "", ""},
		"Hello! , , ",
		HelloHandler,
	},
	{
		[]string{"123", "321"},
		"Hello! 123, 321",
		HelloHandler,
	},
	{
		[]string{"fdsa", "dfas", "asd", "f"},
		"Hello! Fdsa, Dfas, Asd, F",
		HelloHandler,
	},
	{
		[]string{"sasha", "321321"},
		"Hello! Sasha, 321321",
		HelloHandler,
	},
}

func TestPrintHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, test := range tests {
		mockObj := mock.NewMockClientListenerInterface(ctrl)
		mockObj.EXPECT().SendMessage(test.want)
		test.handler(mockObj, test.args)
	}
}
