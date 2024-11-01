package service

import (
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/hafiddna/auth-starterkit-be/dto"
	"github.com/hafiddna/auth-starterkit-be/entity"
	"github.com/hafiddna/auth-starterkit-be/repository"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
)

type StorageService interface {
	Upload(c *fiber.Ctx, createStorageDto dto.CreateStorageDto, file *multipart.FileHeader) (entity.Asset, error)
	FindByPath(getOrDeleteStorageDto dto.GetOrDeleteStorageDto) interface{}
	Delete(asset entity.Asset) error
}

type storageService struct {
	minio           *minio.Client
	assetRepository repository.AssetRepository
}

func NewStorageService(minio *minio.Client, assetRepository repository.AssetRepository) StorageService {
	return &storageService{
		minio:           minio,
		assetRepository: assetRepository,
	}
}

func (s *storageService) Upload(c *fiber.Ctx, createStorageDto dto.CreateStorageDto, file *multipart.FileHeader) (entity.Asset, error) {
	ctx := c.Context()
	contentType := file.Header.Get("Content-Type")
	objectName := file.Filename
	fileSize := file.Size
	fileExtension := strings.Split(objectName, ".")[1]
	unixTime := time.Now().UnixNano()
	newFileName := createStorageDto.Path + "/" + strconv.FormatInt(unixTime, 10) + "." + fileExtension

	src, err := file.Open()
	if err != nil {
		return entity.Asset{}, err
	}

	_, err = s.minio.PutObject(ctx, createStorageDto.Bucket, newFileName, src, fileSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return entity.Asset{}, err
	}

	fileMetadata, err := s.minio.StatObject(ctx, createStorageDto.Bucket, newFileName, minio.StatObjectOptions{})
	jsonFileMetadata, err := json.Marshal(fileMetadata)

	var fileType string
	if strings.Contains(contentType, "image") {
		fileType = "image"
	} else if strings.Contains(contentType, "video") {
		fileType = "video"
	} else if contentType == "application/pdf" {
		fileType = "pdf"
	} else {
		fileType = "file"
	}

	var ownerId string
	var ownerType string
	ownerId, isString := c.Locals("user").(map[string]interface{})["sub"].(string)
	if !isString {
		// TODO: Implement owner type for organization/tenant
		return entity.Asset{}, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	} else {
		ownerType = "users"
	}

	if createStorageDto.Bucket == "private" {
		createStorageDto.Access = "public"
	}
	asset := entity.Asset{
		OwnerID:      &ownerId,
		OwnerType:    &ownerType,
		Name:         objectName,
		Type:         fileType,
		Access:       createStorageDto.Access,
		BucketType:   createStorageDto.Bucket,
		Path:         newFileName,
		Bytes:        float64(fileMetadata.Size),
		FileMetadata: jsonFileMetadata,
	}
	if err = s.assetRepository.Create(asset); err != nil {
		return entity.Asset{}, err
	}

	return asset, nil
}

func (s *storageService) FindByPath(getOrDeleteStorageDto dto.GetOrDeleteStorageDto) interface{} {
	asset := s.assetRepository.FindByPath(getOrDeleteStorageDto.Path)
	if asset.ID == "" {
		return nil
	} else {
		return asset
	}
}

func (s *storageService) Delete(asset entity.Asset) error {
	if err := s.assetRepository.Delete(asset); err != nil {
		return err
	}

	ctx := context.Background()
	err := s.minio.RemoveObject(ctx, asset.BucketType, asset.Path, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}
