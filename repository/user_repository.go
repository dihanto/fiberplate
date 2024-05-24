package repository

import (
	"github.com/dihanto/fiberplate/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

type IUserRepository interface {
	CreateUser(user *model.User) *gorm.DB
}

func (ur *UserRepository) CreateUser(user *model.User) *gorm.DB {
	return ur.db.Model(&model.User{}).Create(user)
}
