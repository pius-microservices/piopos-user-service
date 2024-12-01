package user

import (
	"errors"
	"github.com/pius-microservices/piopos-user-service/package/database/models"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *userRepo {
	return &userRepo{db}
}

func (repo *userRepo) SignUp(data *models.User) (*models.User, error) {

	if err := repo.db.Create(data).Error; err != nil {
		return nil, err
	}

	return data, nil

}

func (repo *userRepo) GetUsers() (*models.Users, error) {
	var data models.Users

	if err := repo.db.
		Select("id, name, username, email, created_at, updated_at").
		Order("created_at DESC").
		Find(&data).Error; err != nil {

		return nil, errors.New("failed to get data")
	}

	if len(data) <= 0 {
		return nil, errors.New("data user is empty")
	}

	return &data, nil
}

func (repo *userRepo) GetUserById(id string) (*models.User, error) {
	var data models.User

	if err := repo.db.
		Preload("Role").
		Find(&data, "id = ?", id).Error; err != nil {
		return nil, errors.New("failed to get data")
	}

	return &data, nil
}

func (repo *userRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	if err := repo.db.
		Where("email = ?", email).
		First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *userRepo) UpdateUser(user *models.User) (*models.User, error) {
	if err := repo.db.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}