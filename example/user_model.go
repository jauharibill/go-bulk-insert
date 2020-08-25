package models

type UserModel struct {
	FullName  string `json:"fullname" gorm:"column:fullname"`
	Age       uint16 `json:"age" gorm:"column:age"`
	BirthDate string `json:"birthdate" gorm:"column:birthdate"`
	MSISDN    string `json:"msisdn" gorm:"column:msisdn"`
}
