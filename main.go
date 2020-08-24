package main

import (
	"go-bulk-insert/models"
	"go-bulk-insert/utils"
)

func main() {
	var user models.UserModel
	utils.ExtractModels(user)
}
