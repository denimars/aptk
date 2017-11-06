package models

type NoUrut struct{
	Id int `gorm:"primary_key;auto_increment;not_null"`
	No int

}