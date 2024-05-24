package app

import (
	"github.com/dihanto/fiberplate/handler"
	"github.com/dihanto/fiberplate/repository"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: fiber.DefaultErrorHandler,
	})
	ur := repository.NewUserRepository(db)
	h := handler.NewHandler(db, *ur)
	h.Routes(app)
	return app
}
