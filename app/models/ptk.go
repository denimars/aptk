package models

type Ptk struct{
	Id int `gorm:"not_null; primary_key; auto_increment"`
	Nupy string `gorm:"size:14"`
	Nama string `gorm:"size:100"`
	TLahir string
	TgLahir string `gorm:"type:date"`
	Alamat string
	Jk string `gorm:"size:2;"`
	StattusPeg string
	GDarah string `gorm:"size:2"`
	NoTelp string `gorm:"size:15"`
	STinggal string
	PTerakhir string
	UKerja string
	SPerkawinan string
	TugasPokok string
	TPokok string
	Tmt string `gorm:"type:date"`
	 
}