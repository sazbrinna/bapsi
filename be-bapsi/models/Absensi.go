package models

import "time"

type Absensi struct {
	IDAbsensi      string    `gorm:"type:varchar(100);primaryKey;unique" json:"id_absensi"` // diganti
	MasukAt        time.Time `gorm:"type:datetime" json:"masuk_at"`
	KeluarAt       time.Time `gorm:"type:datetime" json:"keluar_at"`
	Lokasi         string    `gorm:"type:varchar(32)" json:"lokasi"`
	AkumulasiMenit int       `gorm:"type:integer(11)" json:"akumulasi"`
	IsKeluar       bool      `gorm:"type:tinyint(1)" json:"is_keluar"`

	IDUser string `gorm:"type:varchar(100)" json:"id_user"` // diganti
}
