package di

import (
	"os"

	"github.com/GustavoMS97/go-notes-api/internal/database"
	"github.com/GustavoMS97/go-notes-api/internal/user"
)

type Container struct {
	UserController *user.UserController
}

func NewContainer() *Container {
	mongoClient := database.ConnectMongo()
	dbName := os.Getenv("DATABASE_NAME")
	db := mongoClient.Database(dbName)

	userRepo := user.NewMongoUserRepository(db)
	userService := user.NewUserService(userRepo)
	userController := user.NewUserController(userService)

	return &Container{
		UserController: userController,
	}
}
