// Пакет для создания кастомных ошибок.
package myerrors

import (
	"runtime/debug"
	"time"
)

// Ошибка сохраняющая время создания
func NewErrorTime(errorText string) error {
	return &ErrorWithTime{
		Text: errorText,
		Time: time.Now(),
	}
}

// Ошибка, хранящая стек вызовов, которые привели к ее созникновнию.
func NewErrorTrace(errorText string) error {
	return &ErrorWithStackTrace{
		Text:       errorText,
		StackTrace: debug.Stack(),
	}
}
