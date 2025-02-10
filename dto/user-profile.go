package dto

type UserProfileDTO struct {
	Username        string `json:"username"`
	EmailVerifiedAt int64  `json:"email_verified_at"`
	PhoneVerifiedAt int64  `json:"phone_verified_at"`
	FullName        string `json:"full_name"`
	NickName        string `json:"nick_name"`
	// TODO: Add avatar here
	// TODO: Add list of the teams here
}
