package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	FirstName string `json:"firstName" vaidate:"required"`
	LastName  string `json:"lastName" vaidate:"required"`
}
