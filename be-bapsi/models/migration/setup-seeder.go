package migration

import (
	"fmt"
	"time"

	"github.com/sazbrinna/be_bms/controller/auth"
	"github.com/sazbrinna/be_bms/controller/utils"
	"github.com/sazbrinna/be_bms/controller/utils/id"
	"github.com/sazbrinna/be_bms/models"
)

func DBSeeder() {
	fmt.Println("Processing DBSeeder...")
	var seeder models.Users
	if err := models.DB.First(&seeder).Error; err == nil {
		fmt.Println("Database Already Seeding...")
		return
	}

	// ----------------------------------------------------------------
	// Seed Jam Kerja
	// ----------------------------------------------------------------

	SeedJamKerja("Senin", "09.00", "16.00")
	SeedJamKerja("Selasa", "09.00", "16.00")
	SeedJamKerja("Rabu", "09.00", "16.00")
	SeedJamKerja("Kamis", "09.00", "16.00")
	SeedJamKerja("Jumat", "09.00", "16.00")
	SeedJamKerja("Sabtu", "09.00", "12.00")

	// ----------------------------------------------------------------
	// Seed for divisi
	// ----------------------------------------------------------------

	SeedDivisi("Media Information Center", "asisten")
	SeedDivisi("System Development Center", "asisten")
	SeedDivisi("Admin", "admin")
	SeedDivisi("Staff", "staff")

	// ----------------------------------------------------------------
	// Seed for region
	// ----------------------------------------------------------------

	SeedRegion("Rumah Tanjung", "-6.2415644", "107.0185238")
	SeedRegion("Kampus D", "-6.3705715", "106.8296488")
	SeedRegion("Kampus J1", "-6.248508", "106.970381")

	// ----------------------------------------------------------------
	// Seed for first admin
	// ----------------------------------------------------------------

	userEmail := "asistenbapsi@gmail.com"
	userDivisi := "Admin"
	userJadwal := []string{"Senin", "Rabu", "Sabtu"}

	SeedUserAccount(userEmail, "passbapsi", "admin")
	SeedProfile(userEmail, "admin", userDivisi, "Kampus D", userJadwal)

	// ----------------------------------------------------------------
	// Seed for first admin
	// ----------------------------------------------------------------

	SeedInvitation("zunarvy@gmail.com", "System Development Center", "asisten", "Kampus D")
	SeedInvitation("faizalmaulana.bapsi@gmail.com", "System Development Center", "asisten", "Kampus D")
	SeedInvitation("daniswaramic@gmail.com", "Media Information Center", "asisten", "Kampus D")
	SeedInvitation("sabrinaazka.bapsi@gmail.com", "Media Information Center", "asisten", "Kampus D")
	SeedInvitation("dinasafirabapsi@gmail.com", "Media Information Center", "asisten", "Kampus D")
	SeedInvitation("ibralikmanbapsi@gmail.com", "Media Information Center", "asisten", "Kampus D")
	SeedInvitation("sopyanbapsi@gmail.com", "Media Information Center", "asisten", "Kampus J1")
	SeedInvitation("dewisafirabapsi@gmail.com", "Media Information Center", "asisten", "Kampus J1")

	// Print status ended
	fmt.Println("\nSuccess Seeding")
}

func DayFormat() {
	// ----------------------------------------------------------------
	// Tester dayt format
	// ----------------------------------------------------------------

	year, month, day := time.Now().Date()
	fmt.Println(year, month, day)

	// Convert month to int
	monthInt := int(month)

	tanggal := fmt.Sprintf("%d-%02d-%d", day, monthInt, year)
	fmt.Println(tanggal)

	date := tanggal
	t, err := time.Parse("02-01-2006", date)
	if err != nil {
		panic(err)
	}
	fmt.Println(t.Weekday())
}

func QueryJamKerja(Hari string) models.JamKerja {
	var jamKerja models.JamKerja
	if err := models.DB.Where("hari = ?", utils.Encrypt(Hari)).First(&jamKerja).Error; err != nil {
		fmt.Println("Jam Kerja Tidak Ditemukan")
		return models.JamKerja{}
	}

	return jamKerja
}

func SeedDivisi(divisi, asisten string) {
	dataDivisi := models.Divisi{
		IDDivisi: id.NewV6().String(),
		Divisi:   utils.Encrypt(divisi),
		Group:    utils.Encrypt(asisten),
	}

	models.DB.Create(&dataDivisi)
}

func SeedJamKerja(hari, masuk, pulang string) {
	dataJamKerja := models.JamKerja{
		IDJamKerja: id.NewV6().String(),
		Hari:       utils.Encrypt(hari),
		JamMasuk:   utils.Encrypt(masuk),
		JamPulang:  utils.Encrypt(pulang),
	}

	models.DB.Create(&dataJamKerja)
}

func SeedProfile(email, name, divisi, region string, JadwalMasuk []string) {
	var user models.Users
	if err := models.DB.Where("email = ?", utils.Encrypt(email)).First(&user).Error; err != nil {
		fmt.Println("Email tidak ditemukan")
		return
	}

	var divisis models.Divisi
	if err := models.DB.Where("divisi = ?", utils.Encrypt(divisi)).First(&divisis).Error; err != nil {
		fmt.Println("Divisi tidak ditemukan")
		return
	}

	var regions models.Region
	if err := models.DB.Where("region = ?", utils.Encrypt(region)).First(&regions).Error; err != nil {
		fmt.Println("Region tidak ditemukan")
		return
	}

	dataProfile := models.Profile{
		IDProfile: id.NewV6().String(),
		IDUser:    user.IDUser,
		Nama:      utils.Encrypt(name),
		IDDivisi:  divisis.IDDivisi,
		IDRegion:  regions.IDRegion,
	}

	for _, Jadwal := range JadwalMasuk {
		SaveJadwal := QueryJamKerja(Jadwal)
		dataProfile.JamKerja = append(dataProfile.JamKerja, SaveJadwal)
	}

	models.DB.Create(&dataProfile)
}

func SeedRegion(region, lat, long string) {
	dataRegion := models.Region{
		IDRegion:  id.NewV6().String(),
		Region:    utils.Encrypt(region),
		Latitude:  utils.Encrypt(lat),
		Longitude: utils.Encrypt(long),
	}

	models.DB.Create(&dataRegion)
}

func SeedInvitation(email, divisi, role, region string) {
	dataInvitation := models.Invitation{
		IDInvitation: id.NewV6().String(),
		Email:        utils.Encrypt(email),
		Divisi:       utils.Encrypt(divisi),
		Role:         utils.Encrypt(role),
		Region:       utils.Encrypt(region),
	}

	models.DB.Create(&dataInvitation)
}

func SeedUserAccount(email, password, role string) {
	dataUser := models.Users{
		IDUser:    id.NewV6().String(),
		Email:     utils.Encrypt(email),
		Password:  auth.HashPassword(password),
		Role:      utils.Encrypt(role),
		IsDeleted: false,
	}

	// dataProfile := models.Profile{
	// 	IDProfile: id.NewV6().String(),
	// 	IDDivisi:  divisi.IDDivisi,
	// 	IDUser:    dataUser.IDUser,
	// }

	models.DB.Create(&dataUser)
	// models.DB.Create(&dataProfile)
}
