package finder

import (
	"os"
)

func ValidateDirPath(path string) error {
	p, err := os.Stat(path)
	if err != nil {
		return ErrorPathDoesNotExist
	}

	if !p.IsDir() {
		return ErrorPathIsNotDir
	}

	return nil
}

func ValidateFilePath(fileName string) error {
	f, err := os.Stat(fileName)
	if err != nil {
		return ErrorFileDoesNotExist
	}

	if f.IsDir() {
		return ErrorFileIsDir
	}

	return nil
}
