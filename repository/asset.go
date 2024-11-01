package repository

import (
	"github.com/hafiddna/auth-starterkit-be/entity"
	"gorm.io/gorm"
)

type AssetRepository interface {
	Create(asset entity.Asset) error
	FindByPath(path string) entity.Asset
	Delete(asset entity.Asset) error
}

type assetRepository struct {
	db *gorm.DB
}

func NewAssetRepository(db *gorm.DB) AssetRepository {
	return &assetRepository{db: db}
}

func (r *assetRepository) Create(asset entity.Asset) error {
	return r.db.Create(&asset).Error
}

func (r *assetRepository) FindByPath(path string) entity.Asset {
	var asset entity.Asset
	r.db.Where("path = ?", path).First(&asset)
	return asset
}

func (r *assetRepository) Delete(asset entity.Asset) error {
	return r.db.Delete(&asset).Error
}
