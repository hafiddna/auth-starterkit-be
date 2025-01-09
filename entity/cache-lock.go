package entity

import "github.com/hafiddna/auth-starterkit-be/entity/global"

type CacheLock struct {
	global.Model
	Key        string `gorm:"type:varchar(255);unique_index" json:"key"`
	Owner      string `gorm:"type:uuid" json:"owner"`
	Expiration int64  `gorm:"type:bigint" json:"expiration"`
}

func (CacheLock) TableName() string {
	return "cache_locks"
}
