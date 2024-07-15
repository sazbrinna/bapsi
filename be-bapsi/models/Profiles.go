package models

import "time"

type Profile struct {
	IDProfile string    `gorm:"type:varchar(100);primaryKey;unique" json:"id_profile"` // diganti
	Nama      string    `gorm:"type:varchar(180)" json:"nama"`                         // diganti
	Img       string    `gorm:"type:varchar(50)" json:"img"`
	IsDeleted bool      `gorm:"type:tinyint(1)" json:"is_deleted"`
	CreatedAt time.Time `gorm:"type:datetime" json:"created_at"`

	// foreign keys
	IDUser           string             `gorm:"type:varchar(3)" json:"id_user"`
	IDDivisi         string             `gorm:"type:varchar(3)" json:"id_divisi"`
	IDRegion         string             `gorm:"type:varchar(3)" json:"id_region"`
	JamKerja         []JamKerja         `gorm:"many2many:Profile_JamKerja;"  json:"jam_kerja"`
	AkumulasiBulanan []AkumulasiBulanan `gorm:"foreignKey:IDProfile"  json:"akumulasi_bulanan"`
}
