package model

import "database/sql"

type SecurityQuestion struct {
	Model
	QuestionID string `gorm:"type:uuid;not null;index" json:"question_id"`
	//Question
	DescriptionID sql.NullString `gorm:"type:uuid;index" json:"description_id"`
	//Description
}

func (s *SecurityQuestion) TableName() string {
	return "security_questions"
}
