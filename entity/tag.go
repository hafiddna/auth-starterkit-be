package entity

import "github.com/hafiddna/auth-starterkit-be/entity/global"

type Tag struct {
	global.Model
	Name   string   `gorm:"type:varchar(255);unique" json:"name"`
	Assets []*Asset `gorm:"many2many:asset_tag" json:"assets"`
}
