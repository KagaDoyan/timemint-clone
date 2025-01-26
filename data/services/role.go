package services

import (
	"go-fiber/data/repositories"
	"go-fiber/domain/entities"
	"go-fiber/domain/models"
)

type RoleService interface {
	FindAll(page, limit int) ([]models.Role, int64, error)
	FindByID(id uint) (*models.Role, error)
	Create(role models.Role) (*models.Role, error)
	Update(id uint, role models.Role) (*models.Role, error)
	Delete(id uint) error
}

type roleService struct {
	roleRepo repositories.RoleRepository
}

func NewRoleService(roleRepo repositories.RoleRepository) RoleService {
	return &roleService{
		roleRepo: roleRepo,
	}
}

func (s roleService) FindAll(page, limit int) ([]models.Role, int64, error) {
	data, total, err := s.roleRepo.FindAll(page, limit)
	if err != nil {
		return nil, 0, err
	}
	var result []models.Role
	for _, role := range data {
		result = append(result, models.Role{
			ID:   role.ID,
			Name: role.Name,
		})
	}
	return result, total, nil
}

func (s roleService) FindByID(id uint) (*models.Role, error) {
	role, err := s.roleRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return &models.Role{
		ID:   role.ID,
		Name: role.Name,
	}, nil
}

func (s roleService) Create(role models.Role) (*models.Role, error) {
	roleEntity := entities.Role{
		Name: role.Name,
	}
	data, err := s.roleRepo.Create(roleEntity)
	if err != nil {
		return nil, err
	}
	return &models.Role{
		ID:   data.ID,
		Name: data.Name,
	}, nil
}

func (s roleService) Update(id uint, role models.Role) (*models.Role, error) {
	roleEntity := entities.Role{
		Name: role.Name,
	}
	data, err := s.roleRepo.Update(id, roleEntity)
	if err != nil {
		return nil, err
	}
	return &models.Role{
		ID:   data.ID,
		Name: data.Name,
	}, nil
}

func (s roleService) Delete(id uint) error {
	err := s.roleRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
