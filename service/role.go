package service

import "github.com/hafiddna/auth-starterkit-be/repository"

type RoleService interface {
}

type roleService struct {
	roleRepository repository.RoleRepository
}

func NewRoleService(roleRepository repository.RoleRepository) RoleService {
	return &roleService{
		roleRepository: roleRepository,
	}
}
