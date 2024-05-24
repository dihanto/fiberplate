package handler

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/dihanto/fiberplate/types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"githuh.dom/dihanto/fiberplate/view"
)

func (h *Handler) VRegister(c *fiber.Ctx) error {
	return adaptor.HTTPHandler(templ.Handler(view.RegisterPage()))(c)
}

func (h *Handler) VRegisterPost(c *fiber.Ctx) error {
	remoteData := &types.UserDTOBody{}
	if err := ParseBody(c, remoteData); err != nil {
		return err
	}

	c.Response().Header.Set("HX-Redirect", "/")

	return c.Status(http.StatusOK).SendString("")
}
