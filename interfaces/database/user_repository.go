package database

import (
	"github.com/dionomusuko/go-crean-archi/domain"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (uR *UserRepository) store(user *domain.User) error {
	return uR.db.Save(user).Error
}

func (uR *UserRepository) FindByID(id int) (*domain.User, error) {
	user := domain.User{ID: id}
	err := uR.db.First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (uR *UserRepository) Update(user *domain.User) error {
	return uR.db.Model(&domain.User{ID: user.ID}).Updates(user).Error
}

func (uR *UserRepository) Delete(user *domain.User) error {
	return uR.db.Delete(&user).Error
}

func (uR *UserRepository) FindAll() ([]*domain.User, error) {
	users := []*domain.User{}

	err := uR.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
