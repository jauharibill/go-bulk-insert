package models

type UserModel struct {
	FullName  string `json:"fullname"`
	Age       uint16 `json:"age"`
	BirthDate string `json:"birthdate"`
	MSISDN    string `json:"msisdn"`
}
