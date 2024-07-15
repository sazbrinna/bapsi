package models

type JamKerja struct {
	IDJamKerja string `gorm:"type:varchar(100);primaryKey;unique" json:"id_jam_masuk"` // diganti
	Hari       string `gorm:"type:varchar(36)" json:"hari"`                            // diganti
	JamMasuk   string `gorm:"type:varchar(6)" json:"jam_masuk"`
	JamPulang  string `gorm:"type:varchar(6)" json:"jam_pulang"`

	// foreign keys
	Profile   []Profile `gorm:"many2many:Profile_JamKerja;"  json:"profile"`
	IsDeleted bool      `gorm:"type:tinyint(1)" json:"is_deleted"`
}
