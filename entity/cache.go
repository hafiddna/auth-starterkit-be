package entity

import "github.com/hafiddna/auth-starterkit-be/entity/global"

type Cache struct {
	global.Model
	Key        string `gorm:"type:varchar(255);unique_index" json:"key"`
	Value      string `gorm:"type:text" json:"value"`
	Expiration int64  `gorm:"type:bigint" json:"expiration"`
}

func (c *Cache) TableName() string {
	return "cache"
}
