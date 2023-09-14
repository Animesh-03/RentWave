package repositories

import (
	"user/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[models.User]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	repository := UserRepository{}
	repository.DB = db
	repository.EnsureInitialized()
	return &repository
}

func (r *UserRepository) GetByUsername(username string) (*models.User, error) {
	var user *models.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	return user, err
}
