package staffs

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/controller/utils"
	"github.com/sazbrinna/be_bms/models"
)

func AdmGetAllStaff(c *gin.Context) {
	var staffs []models.Profile
	if err := models.DB.
		Where("is_deleted = ?", 0).
		Find(&staffs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var returnStaff []gin.H

	for _, staff := range staffs {

		// get region name
		var region models.Region
		if err := models.DB.
			Select("region", "id_region").
			Where("id_region = ? AND is_deleted = ?", staff.IDRegion, 0).
			First(&region).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "region not found"})
			return
		}

		// get division name
		var divisi models.Divisi
		if err := models.DB.
			Select("divisi", "id_divisi").
			Where("id_divisi = ? AND is_deleted = ?", staff.IDDivisi, 0).
			First(&divisi).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "divisi not found"})
			return
		}

		var user models.Users
		if err := models.DB.
			Select("role").
			Where("id_user = ? AND is_deleted = ?", staff.IDUser, 0).
			First(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
			return
		}

		// get akumulasi
		now := time.Now()
		month := now.Month().String()
		year := strconv.Itoa(now.Year())

		var akumulasi models.AkumulasiBulanan
		if err := models.DB.
			Select("akumulasi").
			Where("id_profile = ? AND bulan = ? AND tahun = ? AND is_deleted = ?", staff.IDProfile, utils.Encrypt(month), utils.Encrypt(year), false).
			First(&akumulasi).Error; err != nil {
			if err.Error() != "record not found" {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}

		jam := akumulasi.Akumulasi / 60
		menit := akumulasi.Akumulasi % 60

		dataStaff := gin.H{
			"id_profile": staff.IDProfile,
			"id_user":    staff.IDUser,
			"nama":       utils.Decrypt(staff.Nama),
			"role":       utils.Decrypt(user.Role),
			"divisi":     utils.Decrypt(divisi.Divisi),
			"region":     utils.Decrypt(region.Region),
			"tot_jam":    jam,
			"tot_menit":  menit,
		}

		returnStaff = append(returnStaff, dataStaff)
	}

	c.JSON(http.StatusOK, returnStaff)
}
