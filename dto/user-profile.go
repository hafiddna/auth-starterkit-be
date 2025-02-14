package dto

type TeamUserProfileDTO struct {
}

type UserProfileDTO struct {
	Username        string               `json:"username"`
	EmailVerifiedAt int64                `json:"email_verified_at"`
	PhoneVerifiedAt int64                `json:"phone_verified_at"`
	FullName        string               `json:"full_name"`
	NickName        string               `json:"nick_name"`
	Avatar          string               `json:"avatar"`
	Teams           []TeamUserProfileDTO `json:"teams,omitempty"`
}
