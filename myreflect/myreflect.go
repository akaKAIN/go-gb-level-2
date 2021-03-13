package myreflect

import (
	"errors"
	"fmt"
	"github.com/akaKAIN/go-gb-level-2/myfiles"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"reflect"
)

var (
	ErrorNoStruct = errors.New("struct attr passed to a func is nil")
	ErrorNoMap    = errors.New("map attr passed to a func is nil")
)

// Функция обновления значений в полях структуры по соотвествующим
// ключам переданной "мапы"
func UpdateStruct(user *User, values map[string]interface{}) error {
	switch {
	case user == nil:
		return ErrorNoStruct
	case values == nil:
		return ErrorNoMap
	}

	targetVal := reflect.ValueOf(user).Elem()

	for i := 0; i < targetVal.NumField(); i++ {
		fieldName := targetVal.Type().Field(i).Name
		if newVal, ok := values[fieldName]; ok {
			targetVal.Field(i).Set(reflect.ValueOf(newVal))
		}
	}

	return nil
}

func Analyze(fileName, funcName string) (int, error) {
	fileText, err := myfiles.ReadAllFile(fileName)
	if err != nil {
		return 0, err
	}
	fileSet := token.NewFileSet()

	f, err := parser.ParseFile(fileSet, fileName, fileText, 0)
	if err != nil {
		log.Fatal(err)
	}
	if err = ast.Print(fileSet, f); err != nil {
		log.Fatal(err)
	}
	count := 0
	ast.Inspect(f, func(x ast.Node) bool {
		fmt.Println("Inspect")
		if _, ok := x.(*ast.GoStmt); ok {
			count++
		} else {
			fmt.Println("============")
		}
		return false
	})
	fmt.Println(count)
	return 0, nil
}


