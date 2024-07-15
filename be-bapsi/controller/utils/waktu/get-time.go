package waktu

import (
	"strconv"
	"strings"
	"time"

	"github.com/sazbrinna/be_bms/controller/utils"
	"github.com/sazbrinna/be_bms/models"
)

type Time struct {
	StartOfDay   time.Time
	EndOfDay     time.Time
	JamMasukA    int
	MenitMasukA  int
	JamPulangA   int
	MenitPulangA int
}

func SetJadwal(h *Time) error {
	now := time.Now()
	var jadwal models.JamKerja

	// Query the database for the record
	if err := models.DB.
		Where("hari = ?", utils.Encrypt(strconv.Itoa(now.Day()))).
		First(&jadwal).Error; err != nil {
		return err
	}

	// Decrypt and split jadwal.JamMasuk and jadwal.JamPulang
	decJamMasuk := utils.Decrypt(jadwal.JamMasuk)
	decJamPulang := utils.Decrypt(jadwal.JamPulang)
	jamMasuk := strings.Split(decJamMasuk, ".")
	jamPulang := strings.Split(decJamPulang, ".")

	// Convert string to int and assign to h fields
	if len(jamMasuk) == 2 {
		h.JamMasukA, _ = strconv.Atoi(jamMasuk[0])
		h.MenitMasukA, _ = strconv.Atoi(jamMasuk[1])
	}
	if len(jamPulang) == 2 {
		h.JamPulangA, _ = strconv.Atoi(jamPulang[0])
		h.MenitPulangA, _ = strconv.Atoi(jamPulang[1])
	}

	return nil
}

func (h *Time) StartDay() time.Time {
	return h.StartOfDay
}

func (h *Time) EndDay() time.Time {
	return h.EndOfDay
}

func (h *Time) JamMasuk() int {
	return h.JamMasukA
}

func (h *Time) MenitMasuk() int {
	return h.MenitMasukA
}

func (h *Time) JamPulang() int {
	return h.JamPulangA
}

func (h *Time) MenitPulang() int {
	return h.MenitPulangA
}

func NewTime() *Time {
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.AddDate(0, 0, 1).Add(-time.Nanosecond)

	return &Time{
		StartOfDay:   startOfDay,
		EndOfDay:     endOfDay,
		JamMasukA:    0,
		MenitMasukA:  0,
		JamPulangA:   0,
		MenitPulangA: 0,
	}
}
