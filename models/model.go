package models

type Student struct {
	ID   string `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
