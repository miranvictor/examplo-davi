package rh

import (
	"davifiber/internal/app/rh/handler"
	"github.com/gofiber/fiber/v2"
)

func Router(router fiber.Router) {

	rh := router.Group("/rh")
	rh.Get("/funcionarios/:id", handler.GetOne)
	rh.Get("/funcionarios", handler.GetAll)
	rh.Get("/funcionario", handler.GetFirst)
}
