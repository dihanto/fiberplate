package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) Routes(app *fiber.App) {
	h.MediaRoutes(app)
	app.Post("/register", h.CreateUser)
}

func (h *Handler) MediaRoutes(app *fiber.App) {
	app.Get("/register", h.VRegister)
	app.Post("/register", h.VRegisterPost)
}
