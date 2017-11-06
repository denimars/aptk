package models

import "time"

type Log struct{
	Id int `gorm:"primary_key;not_null;auto_increment"`
	Status string
	Waktu time.Time
}