package myerrors

import (
	"fmt"
	"time"
)

// Структура ошибки, хранящая время ее возникновения.
type ErrorWithTime struct {
	Text string    // Text is the error message what we created
	Time time.Time // Time is the time at which the error was created
}

// Формат вывода ошибки с указанием времени ее создания в фиксированном формате
func (e *ErrorWithTime) Error() string {
	return fmt.Sprintf("%s\t%s", e.Time.Format("2006.01.02--15:04:05"), e.Text)
}

// Структура ошибки, хранящая стек вызовов, в которых возникла ошибка.
type ErrorWithStackTrace struct {
	Text       string // Text is the error message what we created
	StackTrace []byte // stack trace of error
}

// Формат вывода ошибки со стэком вызовов.
func (e *ErrorWithStackTrace) Error() string {
	return fmt.Sprintf("%s\n%s", e.Text, e.StackTrace)
}
