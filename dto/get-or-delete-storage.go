package dto

type GetOrDeleteStorageDto struct {
	Path string `query:"path" validate:"required"`
}
