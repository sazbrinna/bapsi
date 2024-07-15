package models

type Divisi struct {
	IDDivisi string `gorm:"type:varchar(3);primaryKey;unique" json:"id_divisi"`
	Divisi   string `gorm:"type:varchar(72)" json:"divisi"` // diganti
	Group    string `gorm:"type:varchar(36)" json:"group"`  // diganti

	Profile   []Profile `gorm:"foreignKey:IDDivisi"  json:"id_profile"`
	IsDeleted bool      `gorm:"type:tinyint(1)" json:"is_deleted"`
}
