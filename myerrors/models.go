package myerrors

import (
	"fmt"
	"time"
)

// Структура ошибки, хранящая время ее возникновения.
type ErrorWithTime struct {
	text string
	time time.Time
}

func (e *ErrorWithTime) formatTime() string {
	return e.time.Format("2006.01.02--15:04:05")
}

func (e *ErrorWithTime) Error() string {
	return fmt.Sprintf("%s\t%s", e.formatTime(), e.text)
}

// Структура ошибки, хранящая стек вызовов, в которых возникла ошибка.
type ErrorWithStackTrace struct {
	text       string
	stackTrace []byte
}

func (e *ErrorWithStackTrace) Error() string {
	return fmt.Sprintf("%s\n%s", e.text, e.stackTrace)
}
