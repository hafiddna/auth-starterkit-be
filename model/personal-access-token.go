package model

import "database/sql"

type PersonalAccessToken struct {
	Model
	TokenableID   int64          `gorm:"not null;type:bigint" json:"tokenable_id"`
	TokenableType string         `gorm:"not null;type:varchar(255)" json:"tokenable_type"`
	Tokenable     interface{}    `gorm:"-" json:"tokenable"`
	Name          string         `gorm:"not null;type:varchar(255)" json:"name"`
	Token         string         `gorm:"not null;type:varchar(64);unique" json:"token"`
	Abilities     sql.NullString `gorm:"type:text;nullable" json:"abilities"`
	LastUsedAt    sql.NullInt64  `gorm:"type:integer;nullable" json:"last_used_at"`
	ExpiresAt     sql.NullInt64  `gorm:"type:integer;nullable" json:"expires_at"`
}

func (p *PersonalAccessToken) TableName() string {
	return "personal_access_tokens"
}
