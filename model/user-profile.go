package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserProfile struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FullName *string            `json:"full_name" bson:"full_name"`
	NickName *string            `json:"nick_name" bson:"nick_name"`
	UserID   string             `json:"user_id" bson:"user_id"`
	//User     *User              `json:"user" bson:"-"`
	AvatarID *string `json:"avatar_id" bson:"avatar_id"`
	//Avatar   *Asset             `json:"avatar" bson:"-"`
	Metadata EmbedJSON `json:"metadata" bson:"metadata"`

	//ClientID            string `json:"ClientID,omitempty" bson:"ClientID,omitempty"`
	//MemberID            string `json:"MemberID,omitempty" bson:"MemberID,omitempty"`
	//QTag                string `json:"QTag,omitempty" bson:"QTag,omitempty"`
	//FullName            string `json:"FullName,omitempty" bson:"FullName,omitempty"`
	//NameConfirmed       bool   `json:"NameConfirmed,omitempty" bson:"NameConfirmed,omitempty"`
	//NIK                 string `json:"NIK,omitempty" bson:"NIK,omitempty"`
	//NIKConfirmed        bool   `json:"NIKConfirmed,omitempty" bson:"NIKConfirmed,omitempty"`
	//Source              string `json:"Source,omitempty" bson:"Source,omitempty"`
	//UpgradeDate         int64  `json:"UpgradeDate,omitempty" bson:"UpgradeDate,omitempty"`
	//ApprovalUpgradeDate int64  `json:"ApprovalUpgradeDate,omitempty" bson:"ApprovalUpgradeDate,omitempty"`
	//ApprovalUpgradeBy   string `json:"ApprovalUpgradeBy,omitempty" bson:"ApprovalUpgradeBy,omitempty"`
	//PasswordHash        string `json:"PasswordHash,omitempty" bson:"PasswordHash,omitempty"`
	//PasswordSalt        string `json:"PasswordSalt,omitempty" bson:"PasswordSalt,omitempty"`
	//FailedPasswordCount int    `json:"FailedAccessCount,omitempty" bson:"FailedAccessCount,omitempty"`
	//PinHash             string `json:"PinHash,omitempty" bson:"PinHash,omitempty"`
	//PinSalt             string `json:"PinSalt,omitempty" bson:"PinSalt,omitempty"`
	//FailedPinCount      int    `json:"FailedPinCount,omitempty" bson:"FailedPinCount,omitempty"`

	//TwoFactorEnabled    bool   `json:"TwoFactorEnabled,omitempty" bson:"TwoFactorEnabled,omitempty"`
	//SecurityStamp       string `json:"SecurityStamp,omitempty" bson:"SecurityStamp,omitempty"`
	//InstallID           string `json:"InstallID,omitempty" bson:"InstallID,omitempty"`
	//DeviceID            string `json:"DeviceID,omitempty" bson:"DeviceID,omitempty"`
	//DeviceType          string `json:"DeviceType,omitempty" bson:"DeviceType,omitempty"`
	//TemporaryBlocked    bool   `json:"TemporaryBlocked,omitempty" bson:"TemporaryBlocked,omitempty"`
	//DocumentCounter     int    `json:"DocumentCounter,omitempty" bson:"DocumentCounter,omitempty"`
	//Sms1Counter         int    `json:"Sms1Counter,omitempty" bson:"Sms1Counter,omitempty"`
	//Sms2Counter         int    `json:"Sms2Counter,omitempty" bson:"Sms2Counter,omitempty"`
	//Sms3Counter         int    `json:"Sms3Counter,omitempty" bson:"Sms3Counter,omitempty"`
	//Sms4Counter         int    `json:"Sms4Counter,omitempty" bson:"Sms4Counter,omitempty"`
	//Status              string `json:"Status,omitempty" bson:"Status,omitempty"`
	//RegisterDate        int64  `json:"RegisterDate,omitempty" bson:"RegisterDate,omitempty"`
	//LastUpdate          int64  `json:"LastUpdate,omitempty" bson:"LastUpdate,omitempty"`
	//LastUpdatedPhone    int64  `json:"LastUpdatedPhone,omitempty" bson:"LastUpdatedPhone,omitempty"`
	//ClosingDate         int64  `json:"ClosingDate,omitempty" bson:"ClosingDate,omitempty"`
	//KK                  string `json:"KK,omitempty" bson:"KK,omitempty"`
	//McID                string `json:"mdID,omitempty" bson:"mdID,omitempty"`
}
