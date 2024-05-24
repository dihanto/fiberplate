package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) Routes(app *fiber.App) {
	h.ApiRoutes(app)
	app.Get("/", h.VIndex)
	app.Get("/register", h.VRegister)
}

func (h *Handler) ApiRoutes(app *fiber.App) {
	app.Post("/register", h.VRegisterPost)
}
