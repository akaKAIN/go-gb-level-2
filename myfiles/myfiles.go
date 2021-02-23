package myfiles

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var (
	ErrorFileIsExists = errors.New("file already exists")
)

// Создаем и записываем данные в новый файл.
func WriteNewFile(fileName string, data []byte) (err error) {
	if ok, err := FileIsExists(fileName); ok {
		return fmt.Errorf("file is exist: %w", err)
	}

	fi, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer func() {
		if err := fi.Close(); err != nil {
			log.Println(err)
		}
	}()

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return
	}

	return
}

func FileIsExists(fileName string) (bool, error) {
	isExist := true

	_, err := os.Stat(fileName)
	if err != nil {
		isExist = false
	}

	return isExist, err
}

func RemoveFile(fileName string) error {
	if err := os.Remove(fileName); err != nil {
		return err
	}
	return nil
}
