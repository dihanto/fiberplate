package handler

import (
	"github.com/dihanto/fiberplate/repository"
	"gorm.io/gorm"
)

type Handler struct {
	UserRepository repository.UserRepository
	db             *gorm.DB
}

func NewHandler(db *gorm.DB, ur repository.UserRepository) *Handler {
	return &Handler{
		UserRepository: ur,
		db:             db,
	}
}
