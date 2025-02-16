package service

import "github.com/hafiddna/auth-starterkit-be/repository"

type PermissionService interface {
}

type permissionService struct {
	permissionRepository repository.PermissionRepository
}

func NewPermissionService(permissionRepository repository.PermissionRepository) PermissionService {
	return &permissionService{
		permissionRepository: permissionRepository,
	}
}
