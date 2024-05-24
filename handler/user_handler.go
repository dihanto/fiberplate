package handler

import (
	"fmt"

	"github.com/dihanto/fiberplate/model"
	"github.com/dihanto/fiberplate/types"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	remoteData := &types.UserRequest{}
	if err := ParseBodyAndValidate(c, remoteData, *h.validator); err != nil {
		return err
	}

	var user = &model.User{}
	user.New(remoteData.User)
	if err := h.UserRepository.CreateUser(user).Error; err != nil {
		fmt.Println(err)
	}

	return c.JSON("create user success")
}
