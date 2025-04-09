package models

import "gorm.io/gorm"

type User_test struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}
