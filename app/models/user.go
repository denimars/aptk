package models

import "time"

type User struct{
	Id int `gorm:"primary_key;auto_increment;not_null"`
	Name string `gorm:"size:100"`
	UserName string `gorm:"size:50";unique`
	Password []byte
	Waktu time.Time

}