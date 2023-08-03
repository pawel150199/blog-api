package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string
	Firstname string
	Lastname  string
	Password  string
}
