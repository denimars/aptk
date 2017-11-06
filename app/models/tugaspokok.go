package models

type TugasPokok struct{
	Id int `gorm:"primary_key;auto_increment;not_null"`
	List string
}