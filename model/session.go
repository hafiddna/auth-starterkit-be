package model

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"github.com/hafiddna/auth-starterkit-be/helper"
	"github.com/yvasiyarov/php_session_decoder/php_serialize"
)

type PreviousPayload struct {
	URL string `json:"url"`
}

type SessionPayload struct {
	Previous PreviousPayload `json:"_previous"`
}

func (s *SessionPayload) SessionEncode() (data string) {
	json := helper.JSONMarshal(s)
	encoded, _ := php_serialize.Serialize(json)
	return base64.StdEncoding.EncodeToString([]byte(encoded))
}

func (s *SessionPayload) SessionDecode(payload string) (err error) {
	decoded, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		return err
	}

	serialized, err := php_serialize.UnSerialize(string(decoded))
	if err != nil {
		return err
	}

	stringSerialized, isString := serialized.(string)
	if !isString {
		return fmt.Errorf("Session data is not a string")
	}

	err = helper.JSONUnmarshal([]byte(stringSerialized), s)

	return nil
}

type Session struct {
	Model
	UserID sql.NullString `gorm:"type:uuid;index;null" json:"user_id"`
	//User         User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
	IPAddress     sql.NullString `gorm:"type:varchar(45)" json:"ip_address"`
	UserAgent     sql.NullString `gorm:"type:text" json:"user_agent"`
	Payload       string         `gorm:"type:text" json:"payload"`
	LastActivity  int64          `gorm:"type:integer;index" json:"last_activity"`
	AppID         string         `gorm:"type:varchar(255);uniqueIndex" json:"app_id"`
	RememberToken sql.NullString `gorm:"type:varchar(100)" json:"remember_token"`
}

func (s *Session) TableName() string {
	return "sessions"
}
