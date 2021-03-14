package finder

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var (
	ErrorFileDoesNotExist = errors.New("file does not exist")
	ErrorFileIsDir        = errors.New("the passed file is a dir")
	ErrorPathDoesNotExist = errors.New("path does not exist")
	ErrorPathIsNotDir     = errors.New("the passed path is not a dir")
)

// Фукция для поиска совпадений файлов по имени и размеру.
// Получает на вход адрес директории для поиска и имя файла для поиска
func FindCopy(path, fileName string) ([]string, error) {
	if err := ValidateDirPath(path); err != nil {
		return nil, err
	}
	if err := ValidateFilePath(fileName); err != nil {
		return nil, err
	}

	var (
		copyFilePathCh = make(chan string)
		copyList       = make([]string, 0)
		checker        Checker
	)
	checker, err := NewSearchTarget(fileName)
	if err != nil {
		return nil, fmt.Errorf("error of init search target struct: %w", err)
	}
	go func() {
		for {
			select {
			case copyPath, ok := <-copyFilePathCh:
				if ok {
					copyList = append(copyList, copyPath)
				} else {
					break
				}
			case <-time.After(10 * time.Second):
				fmt.Println("Timeout. Channel reader is dead")
				return
			}
		}
	}()

	WalkInDir(path, checker, copyFilePathCh)

	checker.WgWait()
	close(copyFilePathCh)
	return copyList, nil
}

// Функция выполняет чтение содержимого директории.
// Найденные директории рекурсивно вызывают WalkInDir.
// Файлы проверяются на копирование. Если факт соответствия копии исходнику подтверждается, то
// путь к файлу отправляется в канал fileCh.
func WalkInDir(path string, c Checker, fileCh chan<- string) {
	c.WgAdd()
	defer c.WgDone()

	dirList, err := os.ReadDir(path)
	if err != nil {
		return
	}

	for _, dirFile := range dirList {
		filePath, err := filepath.Abs(filepath.Join(path, dirFile.Name()))
		if err != nil {
			continue
		}

		fi, err := dirFile.Info()
		if err != nil {
			continue
		}

		if c.Check(fi, filePath) {
			fileCh <- filePath
		}

		if dirFile.IsDir() {
			go WalkInDir(filePath, c, fileCh)
			continue
		}
	}
}
