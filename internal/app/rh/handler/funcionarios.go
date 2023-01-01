package handler

import (
	"davifiber/internal/app/rh/service"
	"github.com/gofiber/fiber/v2"
)

func GetFirst(c *fiber.Ctx) error {
	first, err := service.QueryFirst()

	if err != nil {
		return fiber.ErrBadRequest
	}

	return c.JSON(first)
}

func GetAll(c *fiber.Ctx) error {
	all, err := service.QueryAll()

	if err != nil {
		return fiber.ErrBadRequest
	}

	return c.JSON(all)
}

func GetOne(c *fiber.Ctx) error {
	Id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	one, err := service.QueryOne(Id)
	if err != nil {
		return fiber.ErrBadRequest
	}

	return c.JSON(one)
}
