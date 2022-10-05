package structs

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
}
