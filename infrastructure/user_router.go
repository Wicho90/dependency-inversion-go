package infrastructure

import "github.com/gofiber/fiber/v2"

func Routers(app *fiber.App) {

	userGroup := app.Group("/users")
	userGroup.Get(":id", UserControllerImpl.FindById)

}
