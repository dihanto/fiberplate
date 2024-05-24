package handler

import (
	"fmt"

	"github.com/dihanto/fiberplate/model"
	"github.com/dihanto/fiberplate/types"
	"github.com/dihanto/fiberplate/utils"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateUser(c *fiber.Ctx, remoteData *types.UserRequest) error {
	if err := ParseBodyAndValidate(c, remoteData, *h.validator); err != nil {
		return err
	}

	var user = &model.User{}

	user.New(utils.Convert(*user, remoteData))
	if err := h.UserRepository.CreateUser(user).Error; err != nil {
		fmt.Println(err)
	}

	return c.JSON("create user success")
}
