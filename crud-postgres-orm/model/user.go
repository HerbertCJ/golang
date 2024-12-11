package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       uint   `gorm:"type:int;primary_key"`
	Username string `gorm:"type:varchar(100)"`
	Email    string `gorm:"type:varchar(100)"`
}
