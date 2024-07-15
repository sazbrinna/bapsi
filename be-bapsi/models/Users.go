package models

import "time"

type Users struct {
	IDUser    string    `gorm:"type:varchar(100);primaryKey;unique" json:"id_user"` // diganti
	Email     string    `gorm:"type:varchar(180)" json:"email"`                     // diganti
	Role      string    `gorm:"type:varchar(72)" json:"role"`                       // diganti
	Password  string    `gorm:"type:varchar(180)" json:"password"`                  // diganti
	IsDeleted bool      `gorm:"type:tinyint(1)" json:"is_deleted"`
	CreatedAt time.Time `gorm:"type:datetime" json:"created_at"`

	Profile Profile   `gorm:"foreignKey:IDUser" json:"profile"`
	Absensi []Absensi `gorm:"foreignKey:IDUser"  json:"absensi"`
}
