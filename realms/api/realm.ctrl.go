package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/realbucksavage/miniauth/database/models"
	"github.com/realbucksavage/miniauth/lib/common"
	"github.com/realbucksavage/miniauth/lib/crypto"
	"net/http"
)

type Realm = models.Realm

func list(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var realms []Realm
	if err := db.Order("id desc").Find(&realms).Error; err != nil {
		return
	}

	length := len(realms)
	response := make([]common.JSON, length, length)

	for i := 0; i < length; i++ {
		response[i] = realms[i].Serialize()
	}

	c.JSON(http.StatusOK, response)
}

func findByName(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	name := c.Param("name")

	var realm Realm
	if err := db.Where("name = ?", name).First(&realm).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "REALM_NOT_FOUND"})
		return
	}

	c.JSON(http.StatusOK, realm.Serialize())
}

func create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	type RealmRequest struct {
		Name        string `json:"name" binding:"required"`
		DisplayName string `json:"display_name"`
	}
	var r RealmRequest
	if err := c.BindJSON(&r); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": err.Error()})
		return
	}

	privateKey, err := crypto.GenerateRSAPrivateKey()
	if err != nil {
		fmt.Printf("Cannot generate private key: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": "RSA_PRIV_KEY_ERR"})
		return
	}

	publicKey, err := crypto.GenerateRSAPublicKey(&privateKey.PublicKey)
	if err != nil {
		fmt.Printf("Cannot generate public key: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": "RSA_PUB_KEY_ERR"})
		return
	}

	privateKeyBytes := crypto.EncodePrivateKeyToPem(privateKey)

	if r.DisplayName == "" {
		r.DisplayName = fmt.Sprintf("Realm %s", r.Name)
	}

	realm := Realm{
		Name:        r.Name,
		DisplayName: r.DisplayName,
		PublicKey:   publicKey,
		PrivateKey:  privateKeyBytes,
	}

	db.NewRecord(realm)
	if err := db.Create(&realm).Error; err != nil {
		fmt.Printf("Cannot create realm: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, realm.Serialize())
}

func remove(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	name := c.Param("name")

	var realm Realm
	if err := db.Where("name = ?", name).First(&realm).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "REALM_NOT_FOUND"})
		return
	}

	if err := db.Delete(&realm).Error; err != nil {
		fmt.Printf("Cannot delete realm: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
