package dto

type LoginDTO struct {
	Credential string `json:"credential" validate:"required"`
	Password   string `json:"password" validate:"required,gte=8"`
	Remember   *bool  `json:"remember" validate:"omitempty,boolean"`
}
