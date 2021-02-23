package myerrors

import (
	"runtime/debug"
	"time"
)

func NewErrorTime(errorText string) error {
	return &ErrorWithTime{
		text: errorText,
		time: time.Now(),
	}
}

func NewErrorTrace(errorText string) error {
	return &ErrorWithStackTrace{
		text:       errorText,
		stackTrace: debug.Stack(),
	}
}
