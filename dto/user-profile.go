package dto

type TeamUserProfileDTO struct {
}

type UserProfileDTO struct {
	Username        string               `json:"username"`
	Email           string               `json:"email"`
	EmailVerifiedAt int64                `json:"email_verified_at"`
	Phone           string               `json:"phone"`
	PhoneVerifiedAt int64                `json:"phone_verified_at"`
	FullName        string               `json:"full_name"`
	NickName        string               `json:"nick_name"`
	Avatar          string               `json:"avatar"`
	Role            string               `json:"role"`
	Teams           []TeamUserProfileDTO `json:"teams,omitempty"`
}
