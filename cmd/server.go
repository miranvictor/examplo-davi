package main

import (
	"davifiber/internal/app/rh"
	"davifiber/internal/database"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	errDb := database.Connect()

	if errDb != nil {
		fmt.Println("Erro ao se conectar ao banco", errDb)
		return
	}

	app := fiber.New()

	rh.Router(app)

	err := app.Listen(":4000")
	if err != nil {
		fmt.Println(err)
		return
	}

}
