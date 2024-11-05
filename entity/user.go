package entity

import (
	"github.com/hafiddna/auth-starterkit-be/entity/global"
)

type User struct {
	global.Model
	Email                  string                `gorm:"unique;nullable" json:"email"`
	EmailVerifiedAt        int64                 `gorm:"type:integer;nullable" json:"email_verified_at"`
	Phone                  string                `gorm:"unique;nullable" json:"phone"`
	PhoneVerifiedAt        int64                 `gorm:"type:integer;nullable" json:"phone_verified_at"`
	Username               string                `gorm:"unique;nullable" json:"username"`
	Password               string                `gorm:"select:false" json:"password"`
	TwoFactorSecret        string                `gorm:"nullable;select:false" json:"two_factor_secret"`
	TwoFactorRecoveryCodes string                `gorm:"nullable;select:false" json:"two_factor_recovery_codes"`
	TwoFactorConfirmedAt   int64                 `gorm:"type:integer;nullable" json:"two_factor_confirmed_at"`
	IsActive               bool                  `gorm:"type:boolean" json:"is_active"`
	Role                   string                `gorm:"type:varchar(255);default:'user'" json:"role"`
	RememberToken          string                `gorm:"nullable;select:false" json:"remember_token"`
	MembersOf              []*Team               `gorm:"many2many:team_user" json:"members_of"`
	Teams                  []*Team               `gorm:"foreignKey:OwnerID" json:"teams"`
	Profile                *UserProfile          `gorm:"-" json:"profile"`
	Setting                *UserSetting          `gorm:"-" json:"setting"`
	Roles                  []*Role               `gorm:"many2many:user_role" json:"roles"`
	PasswordResetTokens    []*PasswordResetToken `gorm:"foreignKey:UserID" json:"password_reset_tokens"`
	Sessions               []*Session            `gorm:"foreignKey:UserID" json:"sessions"`
	AssetShares            []*AssetShare         `gorm:"many2many:asset_share" json:"asset_shares"`
	AssetComments          []*AssetComment       `gorm:"many2many:asset_comment" json:"asset_comments"`
	Folders                []*Folder             `gorm:"polymorphicType:OwnerType;polymorphicId:OwnerID;polymorphicValue:users" json:"folders"`
	Assets                 []*Asset              `gorm:"polymorphicType:OwnerType;polymorphicId:OwnerID;polymorphicValue:users" json:"assets"`
}
