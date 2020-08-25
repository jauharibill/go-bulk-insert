package src

import (
	"fmt"
	"math"
	"reflect"
)

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
						columns += fmt.Sprintf("`%v`, ", current_field)
					} else {
						columns += fmt.Sprintf("`%v`", current_field)
					}
				}
			}
		} else {
			if (i + 1) < reflectValue.NumField() {
				columns += fmt.Sprintf("`%v`, ", current_field)
			} else {
				columns += fmt.Sprintf("`%v`", current_field)
			}
		}
	}
	t.SetColumn(columns)
	return t
}

func (t *BulkInsertStruct) ArrangeValue(models [][]interface{}, limitPerInsert int) (s *BulkInsertStruct) {
	var totalRows int = len(models)
	var totalPartialInsert float64
	var totalPartialValues []int
	var dataPartialValues []string

	if totalRows > limitPerInsert {
		lastValue := totalRows % limitPerInsert
		totalPartialInsert = float64(totalRows) / float64(limitPerInsert)
		if lastValue != 0 {
			totalPartialInsert = math.Ceil(float64(totalPartialInsert))
		} else {
			totalPartialInsert = math.Floor(float64(totalPartialInsert))
		}
		totalPartialValues = make([]int, int(totalPartialInsert))
		for i := 0; i < int(totalPartialInsert); i++ {
			totalPartialValues[i] = limitPerInsert
			if (i+1) == int(totalPartialInsert) && lastValue != 0 {
				totalPartialValues[i] = lastValue
			}
		}
	}

	var totalItem int
	dataPartialValues = make([]string, len(totalPartialValues))
	for index, item := range totalPartialValues {
		var values string
		for indexCols := 0; indexCols < item; indexCols++ {
			var valueRows string
			current_rows := models[totalItem+indexCols]
			for indexRows, rows := range current_rows {
				if (indexRows + 1) < len(current_rows) {
					valueRows += fmt.Sprintf("'%v', ", rows)
				} else {
					valueRows += fmt.Sprintf("'%v'", rows)
				}
			}

			if (indexCols + 1) < item {
				values += fmt.Sprintf("(%v),", valueRows)
			} else {
				values += fmt.Sprintf("(%v)", valueRows)
			}
		}
		dataPartialValues[index] = values
		totalItem += item
	}

	t.SetValues(dataPartialValues)
	return t
}

func (t *BulkInsertStruct) InsertQuery(tableName string) (s *BulkInsertStruct) {
	var query string
	for _, item := range t.GetValues() {
		query += fmt.Sprintf("INSERT INTO `%v` (%v) VALUES %v;\n", tableName, t.GetColumn(), item)
	}
	t.SetQuery(query)
	return t
}
