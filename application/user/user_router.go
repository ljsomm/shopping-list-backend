package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ljsomm/shopping-list-backend/domain/user"
	"github.com/ljsomm/shopping-list-backend/infraestructure/datasource"
)

type UserRouter struct{}

func (userRouter UserRouter) InitController(router fiber.Router, datasource datasource.MongoDatasource) {
	userController := NewController(user.NewUseCase(datasource.GetRepository("user")))
	router.Get("/", userController.Index)
}
