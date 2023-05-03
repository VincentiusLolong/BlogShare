package controllers

import (
	"fmt"
	"mailinglist/infrastructure/models"

	"github.com/gofiber/fiber/v2"
)

func (ct *controller) CreateContent(c *fiber.Ctx) error {
	var usercontents models.Contents
	ct.c = *c
	userid := fmt.Sprintf("%v", c.Locals("id"))

	if err := c.BodyParser(&usercontents); err != nil {
		return ct.status(badr, "Json Input", err.Error())
	}

	if validationErr := v.Struct(&usercontents); validationErr != nil {
		return ct.status(badr, "error", validationErr.Error())
	}

	_, err := ct.service.CreateContent(usercontents, userid)
	if err != nil {
		return ct.status(badr, "error", err.Error())
	}

	return ct.status(okr, "result", "content added")
}
