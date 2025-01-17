package entity

import "github.com/hafiddna/auth-starterkit-be/entity/global"

type Team struct {
	global.Model
	OwnerID      string           `gorm:"type:uuid" json:"owner_id"`
	Owner        User             `gorm:"foreignKey:OwnerID" json:"owner,omitempty"`
	Name         string           `json:"name"`
	IsActive     bool             `gorm:"type:boolean;default:false" json:"is_active"`
	Members      []User           `gorm:"many2many:team_user" json:"members,omitempty"`
	Invitations  []TeamInvitation `gorm:"foreignKey:TeamID" json:"invitations,omitempty"`
	Roles        []Role           `gorm:"foreignKey:TeamID" json:"roles,omitempty"`
	Profile      TeamProfile      `gorm:"-" json:"profile,omitempty"`
	Setting      TeamSetting      `gorm:"-" json:"setting,omitempty"`
	Folders      []Folder         `gorm:"polymorphicType:OwnerType;polymorphicId:OwnerID;polymorphicValue:teams" json:"folders,omitempty"`
	Assets       []Asset          `gorm:"polymorphicType:OwnerType;polymorphicId:OwnerID;polymorphicValue:teams" json:"assets,omitempty"`
	PersonalTeam bool             `gorm:"type:boolean;default:false" json:"personal_team"`
}

func (t *Team) TableName() string {
	return "teams"
}
