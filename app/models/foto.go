package models

type Foto struct{
	Id int `gorm:"auto_increment;not_null;primary_key"`
	Nupy string `gorm:"size:14"`
	Url string
}