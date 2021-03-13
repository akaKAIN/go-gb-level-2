package myreflect

import (
	"errors"
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
