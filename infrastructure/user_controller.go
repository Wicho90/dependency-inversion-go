package infrastructure

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wicho90/dependency-inversion-go/application"
)

type UserController struct {
	service application.UserService
}

func NewUserController(service application.UserService) UserController {
	return UserController{service: service}
}

func (u *UserController) FindById(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.JSON(&fiber.Map{
			"error": "Parametro incorrecto" + id,
		})
	}

	user, err := u.service.FinById(id)

	if err != nil {
		return c.JSON(&fiber.Map{
			"error": err,
		})
	}

	return c.JSON(&fiber.Map{
		"data": user,
	})

}
