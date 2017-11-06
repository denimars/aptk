package models


type Anak struct{
	Id int `gorm:"auto_increment; primary_key; not_null"`
	Nupy string `gorm:"size:14"`
	Nama string `gorm:"size:100"`
}