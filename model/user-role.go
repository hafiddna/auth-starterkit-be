package model

type UserRole struct {
	Model
	UserID string `gorm:"type:uuid;not null;index" json:"user_id"`
	//User   User   `gorm:"foreignKey:UserID" json:"user"`
	RoleID string `gorm:"type:uuid;not null;index" json:"role_id"`
	//Role   Role   `gorm:"foreignKey:RoleID" json:"role"`
}

func (u *UserRole) TableName() string {
	return "user_role"
}
