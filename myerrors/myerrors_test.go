package myerrors

import (
	"testing"
)

func TestNewErrorTime(t *testing.T) {
	message := "error with time"
	err := NewErrorTime(message)

	errStruct, ok := err.(*ErrorWithTime)
	if !ok {
		t.Fatal("Wrong error struct")
	}
	if errStruct.Text != message {
		t.Fatalf("wrong message text. Want: %s, Get: %s\n", message, errStruct.Text)
	}
}

func TestNewErrorTrace(t *testing.T) {
	message := "error with stack trace"
	err := NewErrorTrace(message)

	errStruct, ok := err.(*ErrorWithStackTrace)
	if !ok {
		t.Fatal("Wrong error struct", errStruct)
	}
	if errStruct.Text != message {
		t.Fatalf("wrong message text. Want: %s, Get: %s\n", message, errStruct.Text)
	}
}
