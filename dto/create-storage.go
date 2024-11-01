package dto

type CreateStorageDto struct {
	Bucket string `form:"bucket" validate:"required,oneof=public private"`
	Access string `form:"access" validate:"required_if=Bucket public"`
	Path   string `form:"path" validate:"required"`
}
