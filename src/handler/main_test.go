package handler

import (
	mock_main "github.com/golang/mock/gomock"
	"testing"
)

var tests = []struct {
	args []string
}{
	{
		[]string{"1%s%s%s", "2", "3", "45"},
	},
	{w
		[]string{"1%s%s", "2", "34"},
	},
	{
		[]string{"%s%s%s", "1", "2", "3"},
	},
	{
		[]string{"%s", "12"},
	},
}

func TestPrintHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockObj := mock_main.NewMockClientListenerInterface(ctrl)
	mockObj.EXPECT().SendMessage("12345")
	mockObj.EXPECT().SendMessage("1234")
	mockObj.EXPECT().SendMessage("123")
	mockObj.EXPECT().SendMessage("12")

	for _, test := range tests {
		PrintHandler(mockObj, test.args)
	}
}
