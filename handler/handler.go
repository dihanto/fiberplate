package handler

import (
	"github.com/dihanto/fiberplate/repository"
	"gorm.io/gorm"
)

type Handler struct {
	UserRepository repository.UserRepository
	db             *gorm.DB
	validator      *Validator
}

func NewHandler(db *gorm.DB, ur repository.UserRepository) *Handler {
	v := NewValidator()
	return &Handler{
		UserRepository: ur,
		db:             db,
		validator:      v,
	}
}
