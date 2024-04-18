package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ljsomm/shopping-list-backend/domain/user"
)

type UserController struct {
	userService user.UserInputPort
}

func NewController(userService user.UserInputPort) UserController {
	return UserController{
		userService: userService,
	}
}

func (userController *UserController) Index(context *fiber.Ctx) error {
	return context.
		Status(201).
		JSON(user.User{
			Name: "Lucas Juan",
		})
}
