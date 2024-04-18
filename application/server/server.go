package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ljsomm/shopping-list-backend/application/user"
	"github.com/ljsomm/shopping-list-backend/infraestructure/datasource"
)

type Router interface {
	InitController(router fiber.Router, datasource datasource.MongoDatasource)
}

func StartServer() {
	server := fiber.New()

	routers := map[string]Router{
		"/user": user.UserRouter{},
	}

	database_connection, err := datasource.NewMongoConnection()
	if err != nil {
		log.Fatal("error")
		panic(err)
	}
	for base_path, router := range routers {
		router.InitController(server.Group(base_path), database_connection)
	}
	server.Listen(":8080")
}
