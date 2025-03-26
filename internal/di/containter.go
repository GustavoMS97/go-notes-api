package di

import (
	"os"

	"github.com/GustavoMS97/go-notes-api/internal/database"
	"github.com/GustavoMS97/go-notes-api/internal/note"
	"github.com/GustavoMS97/go-notes-api/internal/user"
)

type Container struct {
	UserController *user.UserController
	NoteController *note.NoteController
}

func NewContainer() *Container {
	mongoClient := database.ConnectMongo()
	dbName := os.Getenv("DATABASE_NAME")
	db := mongoClient.Database(dbName)
	// user
	userRepo := user.NewMongoUserRepository(db)
	userService := user.NewUserService(userRepo)
	userController := user.NewUserController(userService)
	// note
	noteRepo := note.NewMongoNoteRepository(db)
	noteService := note.NewNoteService(noteRepo)
	noteController := note.NewNoteController(noteService)

	return &Container{
		UserController: userController,
		NoteController: noteController,
	}
}
