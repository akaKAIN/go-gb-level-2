package myerrors

import (
	"testing"
)

func NewErrorTimeTest(t *testing.T) {
	message := "error error"
	err := NewErrorTime(message)

	errStruct, ok := err.(*ErrorWithTime)
	if !ok {
		t.Fatal("Wrong error struct")
	}
	if errStruct.Text != message {
		t.Fatalf("wrong message text. Want: %s, Get: %s\n", message, errStruct.Text)
	}
}

func NewErrorTraceTest(t *testing.T) {
	message := "error error"
	err := NewErrorTime(message)

	errStruct, ok := err.(*ErrorWithStackTrace)
	if !ok {
		t.Fatal("Wrong error struct")
	}
	if errStruct.Text != message {
		t.Fatalf("wrong message text. Want: %s, Get: %s\n", message, errStruct.Text)
	}
}
