package models

type Invitation struct {
	IDInvitation string `gorm:"type:varchar(100);primaryKey;unique" json:"id_user"` // diganti
	Email        string `gorm:"type:varchar(100)" json:"email"`                     // diganti
	Role         string `gorm:"type:varchar(36)" json:"role"`                       // diganti
	Divisi       string `gorm:"type:varchar(72)" json:"divisi"`                     // diganti
	Region       string `gorm:"type:varchar(62)" json:"region"`                     // diganti
}
