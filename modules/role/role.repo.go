package role

import (
	"errors"

	"github.com/pius-microservices/piopos-user-service/package/database/models"
	"gorm.io/gorm"
)

type roleRepo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *roleRepo {
	return &roleRepo{db}
}

func (repo *roleRepo) AddRole(data *models.Role) (*models.Role, error) {
	if err := repo.db.Create(data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (repo *roleRepo) GetRoles() (*models.Roles, error) {
	var data models.Roles

	if err := repo.db.
		Select("id, name, created_at, updated_at").
		Find(&data).Error; err != nil {
		return nil, errors.New("failed to get data")
	}

	if len(data) <= 0 {
		return nil, errors.New("data role is empty")
	}

	return &data, nil
}

func (repo *roleRepo) GetRoleById(id string) (*models.Role, error) {
	var data models.Role

	if err := repo.db.
		Select("id, name, created_at, updated_at").
		Find(&data, "id = ?", id).Error; err != nil {
		return nil, errors.New("failed to get data")
	}

	return &data, nil
}

func (repo *roleRepo) DeleteRole(id string) error {
	var data models.Role

	if result := repo.db.Delete(data, "id = ?", id).Error; result != nil {
		return result
	}

	return nil
}
