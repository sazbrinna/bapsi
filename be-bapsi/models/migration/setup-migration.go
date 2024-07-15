package migration

import (
	"fmt"
	"log"

	"github.com/sazbrinna/be_bms/models"
)

func DBMigrate() {
	if err := models.DB.AutoMigrate(
		&models.Absensi{},
		&models.AkumulasiBulanan{},
		&models.Divisi{},
		&models.Invitation{},
		&models.JamKerja{},
		&models.Profile{},
		&models.Region{},
		&models.RequestAccount{},
		&models.Users{},
	); err != nil {
		log.Fatalf("Error migrating database: %s", err.Error())
		return
	}
	fmt.Println("Success Migrating")
}
