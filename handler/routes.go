package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) Routes(app *fiber.App) {
	h.MediaRoutes(app)
	app.Get("/", h.VIndex)
	app.Post("/register", h.CreateUser)
	app.Get("/register", h.VRegister)
}

func (h *Handler) MediaRoutes(app *fiber.App) {
	app.Post("/register", h.VRegisterPost)
}
