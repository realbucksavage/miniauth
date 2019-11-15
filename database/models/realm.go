package models

import (
	"github.com/jinzhu/gorm"
	"github.com/realbucksavage/miniauth/lib/common"
)

type Realm struct {
	gorm.Model

	Name        string
	DisplayName string
	PublicKey   string
}

func (r *Realm) Serialize() common.JSON {
	return common.JSON{
		"id":           r.ID,
		"name":         r.Name,
		"display_name": r.DisplayName,
		"public_key":   r.PublicKey,
	}
}

func (r *Realm) Read(m common.JSON) {
	r.ID = uint(m["id"].(float64))
	r.Name = m["name"].(string)
	r.DisplayName = m["display_name"].(string)
	r.PublicKey = m["public_key"].(string)
}
