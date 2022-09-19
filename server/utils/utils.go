package utils

import (
	"errors"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"reflect"
	"strings"
)

func AssignValueFromPath(path string, object any, callback func(string, any)) error {
	value, err := GetValueFromPath(path, object)
	if err != nil {
		return err
	}

	callback(path, value)
	return nil
}

func GetValueFromPath(path string, object any) (any, error) {
	if len(path) == 0 {
		return nil, errors.New("please pass in a string with at least 1 char")
	}

	paths := strings.Split(path, ".")
	fieldName := cases.Title(language.Und).String(paths[0])

	reflectType := reflect.TypeOf(object)
	reflectValue := reflect.ValueOf(object)

	if _, ok := reflectType.FieldByName(fieldName); ok {
		value := reflectValue.FieldByName(fieldName).Interface()

		if len(paths) > 1 {
			newPath := strings.Join(paths[1:], ".")
			return GetValueFromPath(newPath, value)
		}

		return value, nil
	}

	return nil, errors.New(fmt.Sprintf("Could not find a valid field for field name: %s", fieldName))
}
