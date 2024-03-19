package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}
