package entity

import (
	"gorm.io/datatypes"
	"time"
)

type User struct {
	ID                     string     `gorm:"type:uuid;primaryKey" json:"id"`
	Email                  *string    `gorm:"unique;nullable" json:"email"`
	EmailVerifiedAt        *time.Time `gorm:"type:timestamp;nullable" json:"email_verified_at"`
	Phone                  *string    `gorm:"unique;nullable" json:"phone"`
	PhoneVerifiedAt        *time.Time `gorm:"type:timestamp;nullable" json:"phone_verified_at"`
	Username               *string    `gorm:"unique;nullable" json:"username"`
	Password               string     `gorm:"select:false" json:"password"`
	TwoFactorSecret        *string    `gorm:"nullable;select:false" json:"two_factor_secret"`
	TwoFactorRecoveryCodes *string    `gorm:"nullable;select:false" json:"two_factor_recovery_codes"`
	TwoFactorConfirmedAt   *time.Time `gorm:"type:timestamp;nullable" json:"two_factor_confirmed_at"`
	IsActive               bool       `gorm:"type:boolean" json:"is_active"`
	Role                   string     `gorm:"type:varchar(255)" json:"role"`
	RememberToken          *string    `gorm:"nullable;select:false" json:"remember_token"`
	//Roles                  []Role         `gorm:"many2many:role_user" json:"roles"`
	//Sessions               []Session      `gorm:"foreignKey:UserID" json:"sessions"`
	Metadata datatypes.JSON `gorm:"type:jsonb" json:"metadata"`
}
