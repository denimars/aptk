package models

type TugasTambahan struct{
	Id int `gorm:"primary_key;auto_increment;not_null"`
	Nupy string `gorm:"size:14"`
	Tugas string
}