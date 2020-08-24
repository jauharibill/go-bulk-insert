package utils

import (
	"fmt"
	"reflect"
)

func ExtractModelField(model interface{}, exceptionalField []string) {
	var reflectValue = reflect.ValueOf(model)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		for _, item := range exceptionalField {
			current_field := reflectType.Field(i).Tag.Get("json")
			if current_field != item {
				fmt.Println(current_field)
			}
		}
	}
}

func PutAllValueTogether(models [][]interface{}) string {
	var values string

	for indexCols, cols := range models {
		var valueRows string
		for indexRows, rows := range cols {
			if (indexRows + 1) < len(cols) {
				valueRows += fmt.Sprintf("'%v', ", rows.(string))
			} else {
				valueRows += fmt.Sprintf("'%v' ", rows.(string))
			}
		}

		if (indexCols + 1) < len(cols) {
			values += fmt.Sprintf("(%v),", valueRows)
		} else {
			values += fmt.Sprintf("(%v);", valueRows)
		}
	}

	return values
}
