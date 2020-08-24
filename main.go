package main

import (
	"fmt"
	"go-bulk-insert/models"
	"go-bulk-insert/utils"
)

func main() {
	var user models.UserModel
	bulk := new(utils.BulkInsertStruct)

	value := [][]interface{}{
		{"bill", 12, "23-10-1994", "082245088948"},
		{"bill", 12, "23-10-1994", "082245088948"},
		{"bill", 12, "23-10-1994", "082245088948"},
	}
	fmt.Println(bulk.ArrangeValue(value).
		ExtractModelField(user, []string{}).
		PutQueryTogether("users").
		GetQuery())
}
