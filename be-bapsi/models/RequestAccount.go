package models

import "time"

type RequestAccount struct {
	IDRequest string    `gorm:"type:varchar(100);primaryKey;unique" json:"id_request"` // diganti
	Nama      string    `gorm:"type:varchar(180)" json:"nama"`                         // diganti
	Email     string    `gorm:"type:varchar(180)" json:"email"`                        // diganti
	Role      string    `gorm:"type:varchar(72)" json:"role"`                          // diganti
	IDDivisi  string    `gorm:"type:varchar(100)" json:"id_divisi"`                    // diganti
	IDRegion  string    `gorm:"type:varchar(100)" json:"id_region"`                    // diganti
	Password  string    `gorm:"type:varchar(180)" json:"password"`                     // diganti
	CreatedAt time.Time `gorm:"type:datetime" json:"created_at"`
}
