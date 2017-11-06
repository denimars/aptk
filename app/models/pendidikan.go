package models

type Pendidikan struct{
	Id int `gorm:"primary_key;not_null;auto_increment"`
	List string
}
