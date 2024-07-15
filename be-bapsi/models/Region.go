package models

type Region struct {
	IDRegion  string `gorm:"type:varchar(100);primaryKey;unique" json:"id_region"` // diganti
	Region    string `gorm:"type:varchar(36)" json:"region"`                       // diganti
	Latitude  string `gorm:"type:varchar(15)" json:"latitude"`
	Longitude string `gorm:"type:varchar(15)" json:"longitude"`
	IsDeleted bool   `gorm:"type:tinyint(1)" json:"is_deleted"`

	Profile []Profile `gorm:"foreignKey:IDRegion"  json:"id_profile"`
}
