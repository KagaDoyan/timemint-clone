package repositories

import (
	"go-fiber/domain/entities"

	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

type RoleRepository interface {
	FindAll(page, limit int) ([]entities.Role, int64, error)
	FindByID(id uint) (*entities.Role, error)
	Create(role entities.Role) (*entities.Role, error)
	Update(id uint, role entities.Role) (*entities.Role, error)
	Delete(id uint) error
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	db.AutoMigrate(&entities.Role{})
	return &roleRepository{db: db}
}

func (s roleRepository) FindAll(page, limit int) ([]entities.Role, int64, error) {
	var roles []entities.Role
	var total int64

	// Calculate offset
	offset := (page - 1) * limit

	// Count total records
	if err := s.db.Model(&entities.Role{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Find paginated records
	err := s.db.Offset(offset).Limit(limit).Find(&roles).Error
	if err != nil {
		return nil, 0, err
	}

	return roles, total, nil
}

func (s roleRepository) FindByID(id uint) (*entities.Role, error) {
	var role entities.Role
	err := s.db.First(&role, id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (s roleRepository) Create(role entities.Role) (*entities.Role, error) {
	err := s.db.Create(&role).Error
	if err != nil {
		return nil, err
	}
	var result entities.Role
	err = s.db.First(&result, role.ID).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (s roleRepository) Update(id uint, role entities.Role) (*entities.Role, error) {
	err := s.db.Model(&entities.Role{}).Where("id = ?", id).Updates(&role).Error
	if err != nil {
		return nil, err
	}
	var result entities.Role
	err = s.db.First(&result, id).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s roleRepository) Delete(id uint) error {
	err := s.db.Unscoped().Delete(&entities.Role{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
