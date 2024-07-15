package profile

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sazbrinna/be_bms/controller/utils"
	"github.com/sazbrinna/be_bms/models"
)

func GetProfile(c *gin.Context) {
	// this is for profile page only on his own

	userID, exists := c.Get("id_user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	var profile models.Profile
	if err := models.DB.Select("id_profile", "nama", "id_divisi", "id_region").Where("id_user = ?", userID).First(&profile).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var region models.Region
	if err := models.DB.Select("region").Where("id_region = ?", profile.IDRegion).First(&region).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var divisi models.Divisi
	if err := models.DB.Select("divisi").Where("id_divisi = ?", profile.IDDivisi).First(&divisi).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now()
	month := now.Month().String()
	year := strconv.Itoa(now.Year())

	var akumulasi models.AkumulasiBulanan
	if err := models.DB.
		Select("akumulasi").
		Where("id_profile = ? AND bulan = ? AND tahun = ? AND is_deleted = ?", profile.IDProfile, utils.Encrypt(month), utils.Encrypt(year), false).
		First(&akumulasi).Error; err != nil {
		if err.Error() != "record not found" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	jam := akumulasi.Akumulasi / 60
	menit := akumulasi.Akumulasi % 60

	c.JSON(http.StatusOK, gin.H{
		"msg":             "success",
		"nama":            utils.Decrypt(profile.Nama),
		"divisi":          utils.Decrypt(divisi.Divisi),
		"region":          utils.Decrypt(region.Region),
		"akumulasi_jam":   jam,
		"akumulasi_menit": menit,
	})
}

func GetProfileOther(c *gin.Context) {
	userID := c.Param("id")

	var profile models.Profile
	if err := models.DB.Select("id_user", "id_profile", "nama", "id_divisi", "id_region").Where("id_user = ?", userID).First(&profile).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var region models.Region
	if err := models.DB.Select("region").Where("id_region = ?", profile.IDRegion).First(&region).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var divisi models.Divisi
	if err := models.DB.Select("divisi").Where("id_divisi = ?", profile.IDDivisi).First(&divisi).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now()
	month := now.Month().String()
	year := strconv.Itoa(now.Year())

	var akumulasi models.AkumulasiBulanan
	if err := models.DB.
		Select("akumulasi").
		Where("id_profile = ? AND bulan = ? AND tahun = ? AND is_deleted = ?", profile.IDProfile, utils.Encrypt(month), utils.Encrypt(year), false).
		First(&akumulasi).Error; err != nil {
		if err.Error() != "record not found" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	jam := akumulasi.Akumulasi / 60
	menit := akumulasi.Akumulasi % 60

	c.JSON(http.StatusOK, gin.H{
		"msg":             "success",
		"id_user":         profile.IDUser,
		"nama":            utils.Decrypt(profile.Nama),
		"divisi":          utils.Decrypt(divisi.Divisi),
		"region":          utils.Decrypt(region.Region),
		"akumulasi_jam":   jam,
		"akumulasi_menit": menit,
	})
}
