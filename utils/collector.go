package utils

import (
	"fmt"
	"reflect"
)

type BulkInsertStruct struct {
	Columns *string
	Query   *string
	Values  *string
}

func (t *BulkInsertStruct) ExtractModelField(model interface{}, exceptionalField []string) (s *BulkInsertStruct) {
	var reflectValue = reflect.ValueOf(model)
	var columns string

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		current_field := reflectType.Field(i).Tag.Get("json")
		if len(exceptionalField) > 0 {
			for _, item := range exceptionalField {
				if current_field != item {
					if (i + 1) < reflectValue.NumField() {
						columns += fmt.Sprintf("'%v', ", current_field)
					} else {
						columns += fmt.Sprintf("'%v'", current_field)
					}
				}
			}
		} else {
			if (i + 1) < reflectValue.NumField() {
				columns += fmt.Sprintf("'%v', ", current_field)
			} else {
				columns += fmt.Sprintf("'%v'", current_field)
			}
		}

	}
	t.SetColumn(columns)
	return t
}

func (t *BulkInsertStruct) ArrangeValue(models [][]interface{}) (s *BulkInsertStruct) {
	var values string
	for indexCols, cols := range models {
		var valueRows string
		for indexRows, rows := range cols {
			if (indexRows + 1) < len(cols) {
				valueRows += fmt.Sprintf("'%v', ", rows)
			} else {
				valueRows += fmt.Sprintf("'%v'", rows)
			}
		}

		if (indexCols + 1) < len(models) {
			values += fmt.Sprintf("(%v),", valueRows)
		} else {
			values += fmt.Sprintf("(%v);", valueRows)
		}
	}
	t.SetValues(values)
	return t
}

func (t *BulkInsertStruct) PutQueryTogether(tableName string) (s *BulkInsertStruct) {
	t.SetQuery(fmt.Sprintf("INSERT INTO %v (%v) VALUES %v", tableName, t.GetColumn(), t.GetValues()))
	return t
}
