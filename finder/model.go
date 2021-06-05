package finder

import (
	"errors"
	"go.uber.org/zap"
	"log"
	"os"
	"path/filepath"
	"sync"
)



var (
	ErrorWrongFileName = errors.New("wrong file name")
)

type SearchTarget struct {
	Name string
	Path string
	Size int64
	wg   *sync.WaitGroup
	logger *zap.Logger
}

type Checker interface {
	Check(info os.FileInfo, filePath string) bool
	WgAdd()
	WgDone()
	WgWait()
	ShowError(string)
	ShowInfo(string)
}

func (s *SearchTarget) Check(fi os.FileInfo, filePath string) bool {
	if s.Path == filePath {
		return false
	}
	return s.isCopy(fi)
}

func (s *SearchTarget) isCopy(fi os.FileInfo) bool {
	return s.Name == fi.Name() && s.Size == fi.Size()
}

func (s *SearchTarget) WgAdd() {
	s.wg.Add(1)
}

func (s *SearchTarget) WgDone() {
	s.wg.Done()
}

func (s *SearchTarget) WgWait() {
	s.wg.Wait()
}

func (s *SearchTarget) ShowError(message string) {
	s.logger.Error(message)
}

func (s *SearchTarget) ShowInfo(message string) {
	s.logger.Info(message)
}

func NewSearchTarget(fileName string) (*SearchTarget, error) {
	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			log.Printf("logger init error: %s\n", fileName)
		}
	}(logger)
	logger = logger.With(zap.String("fileName", fileName))

	if fileName == "" {
		return nil, ErrorWrongFileName
	}

	fi, err := os.Stat(fileName)
	if err != nil {
		return nil, err
	}

	path, err := filepath.Abs(fileName)
	if err != nil {
		return nil, err
	}

	return &SearchTarget{
		Name: fi.Name(),
		Path: path,
		Size: fi.Size(),
		wg:   new(sync.WaitGroup),
		logger: logger,
	}, nil
}
