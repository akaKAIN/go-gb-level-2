// Пакет для работы с файловой системой
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
	if FileIsExists(fileName) {
		return ErrorFileIsExists
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

//Функция проверки существования указанного файла
func FileIsExists(fileName string) bool {
	if _, err := os.Stat(fileName); err == nil {
		return true
	}
	return false
}

//Функция удаления указанного файла
func RemovingFile(fileName string) error {
	if err := os.Remove(fileName); err != nil {
		return err
	}
	return nil
}
