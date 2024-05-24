package handler

import (
	"github.com/dihanto/fiberplate/utils"
	"github.com/gofiber/fiber/v2"
)

func ParseBody(c *fiber.Ctx, body interface{}) *utils.RequestError {
	if err := c.BodyParser(body); err != nil {
		return utils.RequestErrorFrom(&utils.BAD_REQUEST, err.Error())
	}

	return nil
}
