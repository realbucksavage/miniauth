package realms

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/realbucksavage/miniauth/database/models"
	"github.com/realbucksavage/miniauth/lib"
	"github.com/realbucksavage/miniauth/lib/crypto"
	"net/http"
)

func initMasterRealm(c *gin.Context) {

	// Only allowed locally
	if !lib.IsLocalRequest(c.Request.RemoteAddr) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	db := c.MustGet("db").(*gorm.DB)

	var masterRealm models.Realm
	db.Where("name = ?", "master").First(&masterRealm)

	// FIXME: There's a better way for sure
	if masterRealm.Name == "master" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "ERR_REALM_EXISTS"})
		return
	}

	type AdminUserDef struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var r AdminUserDef
	if err := c.BindJSON(&r); err != nil {
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

	masterRealm = models.Realm{
		Name:        "master",
		DisplayName: "Master",
		PublicKey:   publicKey,
		PrivateKey:  privateKeyBytes,
	}

	db.NewRecord(masterRealm)
	if err := db.Create(&masterRealm).Error; err != nil {
		fmt.Printf("Cannot create master realm: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "INTERNAL_SERVER_ERR",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, masterRealm.Serialize())
}
