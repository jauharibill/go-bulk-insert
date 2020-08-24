package utils

import (
	"fmt"
	"reflect"
)

func ExtractModels(model interface{}) {
	var reflectValue = reflect.ValueOf(model)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println(reflectType.Field(i).Tag.Get("json"))
	}
}
