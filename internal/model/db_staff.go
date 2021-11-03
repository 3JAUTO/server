package model

import "gorm.io/gorm"

// DBStaff ...
type DBStaff struct {
	gorm.Model
	Email     string
	Phone     string
	Password  string
	Privilege int8
}
