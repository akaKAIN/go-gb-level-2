package myfiles

import (
	"errors"
	"log"
	"os"
)

var (
	ErrorFileIsExists = errors.New("file already exists")
)

// Создаем и записываем данные в новый файл.
func WriteNewFile(fileName string, data []byte) (err error) {
	if _, err = os.Stat(fileName); err == nil {
		err = ErrorFileIsExists
		return
	}

	fi, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer func() {
		if closeErr := fi.Close(); closeErr != nil {
			log.Println(closeErr)
		}
	}()

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return
	}

	return
}
