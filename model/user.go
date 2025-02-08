package model

import "database/sql"

type User struct {
	Model
	Email                  sql.NullString `gorm:"unique;nullable" json:"email"`
	EmailVerifiedAt        sql.NullInt64  `gorm:"type:integer;nullable" json:"email_verified_at"`
	Phone                  sql.NullString `gorm:"unique;nullable" json:"phone"`
	PhoneVerifiedAt        sql.NullInt64  `gorm:"type:integer;nullable" json:"phone_verified_at"`
	Username               sql.NullString `gorm:"unique;nullable" json:"username"`
	Password               string         `gorm:"select:false" json:"password"`
	Pin                    sql.NullString `gorm:"select:false" json:"pin"`
	TwoFactorSecret        sql.NullString `gorm:"nullable;select:false" json:"two_factor_secret"`
	TwoFactorRecoveryCodes sql.NullString `gorm:"nullable;select:false" json:"two_factor_recovery_codes"`
	TwoFactorConfirmedAt   sql.NullInt64  `gorm:"type:integer;nullable" json:"two_factor_confirmed_at"`
	IsActive               bool           `gorm:"type:boolean" json:"is_active"`
	RememberToken          sql.NullString `gorm:"nullable;select:false" json:"remember_token"`
	//MembersOf              []Team               `gorm:"many2many:team_user" json:"members_of"`
	//Teams                  []Team               `gorm:"foreignKey:OwnerID" json:"teams"`
	//Profile                UserProfile          `gorm:"-" json:"profile"`
	//Setting                UserSetting          `gorm:"-" json:"setting"`
	Roles []Role `gorm:"many2many:user_role;joinForeignKey:UserID;joinReferences:RoleID" json:"roles"`
	//PasswordResetTokens    []PasswordResetToken `gorm:"foreignKey:UserID" json:"password_reset_tokens"`
	//Sessions               []Session            `gorm:"foreignKey:UserID" json:"sessions"`
	//AssetShares            []AssetShare         `gorm:"many2many:asset_share" json:"asset_shares"`
	//AssetComments          []AssetComment       `gorm:"many2many:asset_comment" json:"asset_comments"`
	//Folders                []Folder             `gorm:"polymorphicType:OwnerType;polymorphicId:OwnerID;polymorphicValue:users" json:"folders"`
	//Assets                 []Asset              `gorm:"polymorphicType:OwnerType;polymorphicId:OwnerID;polymorphicValue:users" json:"assets"`
	//SecurityQuestions      []SecurityQuestion   `gorm:"many2many:user_security_answers" json:"security_questions"`
}

func (u *User) TableName() string {
	return "users"
}
