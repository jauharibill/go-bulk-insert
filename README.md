# GORM BULK INSERT

## Introduction

this library is used for creating raw query insert in bulk mode. this library created because gorm library doesnt support bulk request, so this library created for supporting gorm library only.


### How to install

```
go get -u github.com/jauharibill/gobulkinsert
```

### How to use

```
package main

import (
	"fmt"

	"github.com/jauharibill/gobulkinsert"
)

type UserModel struct {
	FullName string `json:"fullname"`
	Age      string `json:"age"`
	Phone    string `json:"phone"`
}

func main() {
	var user UserModel

	value := [][]interface{}{
		{"bill", 12, "17-09-1945", "082245xxxx"},
		{"bill", 12, "17-09-1945", "082245xxxx"},
		{"bill", 12, "17-09-1945", "082245xxxx"},
		{"bill", 12, "17-09-1945", "082245xxxx"},
		{"bill", 12, "17-09-1945", "082245xxxx"},
		{"bill", 12, "17-09-1945", "082245xxxx"},
		{"bill", 12, "17-09-1945", "082245xxxx"},
		{"bill", 12, "17-09-1945", "082245xxxx"},
		{"bill", 12, "17-09-1945", "082245xxxx"},
		{"bill", 12, "17-09-1945", "082245xxxx"},
		{"bill", 12, "17-09-1945", "082245xxxx"},
		{"bill", 12, "17-09-1945", "082245xxxx"},
		{"bill", 12, "17-09-1945", "082245xxxx"},
		{"bill", 12, "17-09-1945", "082245xxxx"},
		{"bill", 12, "17-09-1945", "082245xxxx"},
	}
  
	query := gobulkinsert.Service().
		ExtractModelField(user, []string{}).
		ArrangeValue(value, 6).
		InsertQuery("users").
		GetQuery()

	fmt.Println(query)
}

```

### Output

```
INSERT INTO `users` (`fullname`, `age`, `phone`) VALUES ('bill', '12', '17-09-1945', '082245xxxx'),('bill', '12', '17-09-1945', '082245xxxx'),('bill', '12', '17-09-1945', '082245xxxx'),('bill', '12', '17-09-1945', '082245xxxx'),('bill', '12', '17-09-1945', '082245xxxx'),('bill', '12', '17-09-1945', '082245xxxx');
INSERT INTO `users` (`fullname`, `age`, `phone`) VALUES ('bill', '12', '17-09-1945', '082245xxxx'),('bill', '12', '17-09-1945', '082245xxxx'),('bill', '12', '17-09-1945', '082245xxxx'),('bill', '12', '17-09-1945', '082245xxxx'),('bill', '12', '17-09-1945', '082245xxxx'),('bill', '12', '17-09-1945', '082245xxxx');
INSERT INTO `users` (`fullname`, `age`, `phone`) VALUES ('bill', '12', '17-09-1945', '082245xxxx'),('bill', '12', '17-09-1945', '082245xxxx'),('bill', '12', '17-09-1945', '082245xxxx');
```

### How to implement in GORM Library

put generated query in variable, and then execute it with `Exec()` method in query builder gorm. for example :

```
func (i *gorm.DB) test () {
  query := <this value belongs to gobulkrequest>
  i.Exec(query) // you could put any db transaction also
}
```

## Contribute

Please pull request if you have any improvement, Cheers!
