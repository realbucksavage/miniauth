package models

import (
	"github.com/jinzhu/gorm"
	"github.com/realbucksavage/miniauth/lib/common"
)

type Realm struct {
	gorm.Model

	Name        string `gorm:"unique;index:idx_realm_name"`
	DisplayName string
	PublicKey   []byte `gorm:"not null"`
	PrivateKey  []byte `gorm:"not null"`
}

func (r *Realm) Serialize() common.JSON {
	return common.JSON{
		"name":         r.Name,
		"display_name": r.DisplayName,
		"public_key":   string(r.PublicKey),
	}
}
