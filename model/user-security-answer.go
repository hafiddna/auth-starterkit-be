package model

type UserSecurityAnswer struct {
	Model
	UserID             string `gorm:"type:uuid;not null;index" json:"user_id"`
	SecurityQuestionID string `gorm:"type:uuid;not null;index" json:"security_question_id"`
	Answer             string `gorm:"type:text;not null" json:"answer"`
}

func (s *UserSecurityAnswer) TableName() string {
	return "user_security_answers"
}
