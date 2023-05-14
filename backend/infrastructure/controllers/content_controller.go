package controllers

import (
	"errors"
	"fmt"
	"mailinglist/infrastructure/models"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
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

func (ct *controller) GetAccountContent(c *fiber.Ctx) error {
	userid := fmt.Sprintf("%v", c.Locals("id"))
	contentuser, errs := ct.service.GetUserContent(userid)
	if errs != nil {
		if errors.Is(errs, pgx.ErrNoRows) {
			return &fiber.Error{Code: fiber.StatusNotFound, Message: "Content Not Found"}
		} else {
			return &fiber.Error{Code: fiber.StatusInternalServerError, Message: "Failed to get content"}
		}
	}
	return c.JSON(contentuser)
}

func (ct *controller) EditContent(c *fiber.Ctx) error {
	userid := fmt.Sprintf("%v", c.Locals("id"))

	var usercontents models.GetContent

	if err := c.BodyParser(&usercontents); err != nil {
		return &fiber.Error{Code: fiber.StatusNotFound, Message: err.Error()}
	}

	if validationErr := v.Struct(&usercontents); validationErr != nil {
		return &fiber.Error{Code: fiber.StatusNotFound, Message: validationErr.Error()}
	}

	updateduser, errs := ct.service.EditContent(usercontents, userid)
	if errs != nil {
		return &fiber.Error{Code: fiber.StatusNotFound, Message: "Cannot Edit"}
	}

	return c.JSON(updateduser)
}
