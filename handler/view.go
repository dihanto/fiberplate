package handler

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/dihanto/fiberplate/types"
	"github.com/dihanto/fiberplate/view"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func (h *Handler) VIndex(c *fiber.Ctx) error {
	return adaptor.HTTPHandler(templ.Handler(view.IndexPage()))(c)
}

func (h *Handler) VRegister(c *fiber.Ctx) error {
	return adaptor.HTTPHandler(templ.Handler(view.RegisterPage()))(c)
}

func (h *Handler) VRegisterPost(c *fiber.Ctx) error {
	remoteData := &types.UserRequest{}

	h.CreateUser(c, remoteData)

	c.Response().Header.Set("HX-Redirect", "/")

	return c.Status(http.StatusOK).SendString("")
}
