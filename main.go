package main

import (
	"go-bulk-insert/models"
	"go-bulk-insert/utils"
)

func main() {
	var user models.UserModel
	var exceptField = []string{"fullname"}
	utils.ExtractModelField(user, exceptField)
}
