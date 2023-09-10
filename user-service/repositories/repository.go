package repositories

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository[T any] struct {
	DB *gorm.DB
}

func (r *Repository[T]) EnsureInitialized() {
	var entity *T
	if !r.DB.Migrator().HasTable(&entity) {
		r.DB.Migrator().CreateTable(&entity)
	}
}

// Get all the entities in the table
func (r *Repository[T]) All() (*[]T, error) {
	var entities *[]T
	err := r.DB.Find(&entities).Error

	if err != nil {
		fmt.Println("Error adding entity: ", err)
		return nil, err
	}

	return entities, nil
}

// Add a single entity
func (r *Repository[T]) AddEntity(entity *T) error {
	return r.DB.Create(&entity).Error
}

// Add multiple entities
func (r *Repository[T]) AddEntities(entity *[]T) error {
	return r.DB.Create(&entity).Error
}

// Get Entity by the Primary Key
func (r *Repository[T]) GetByID(id uint) (*T, error) {
	var entity *T

	err := r.DB.First(&entity, id).Error

	if err != nil {
		fmt.Println("Error adding entity: ", err)
		return nil, err
	}

	return entity, nil
}
