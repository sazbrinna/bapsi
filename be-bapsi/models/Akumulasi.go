package models

type AkumulasiBulanan struct {
	IDAkumulasiBulanan string `gorm:"type:varchar(100);primaryKey;unique" json:"id_akumulasi_bulanan"` // diganti
	Tahun              string `gorm:"type:varchar(4)" json:"tahun"`
	Bulan              string `gorm:"type:varchar(10)" json:"bulan"`
	Akumulasi          int    `gorm:"type:varchar(255)" json:"akumulasi"`
	IsDeleted          bool   `gorm:"type:tinyint(1)" json:"is_deleted"`

	// foreign keys,,,n
	IDProfile string `gorm:"type:varchar(100)" json:"id_profile"` // diganti
}
