package repository

import (
	"errors"

	"github.com/GustavoMS97/go-notes-api/internal/user/entity"
)

type UserRepository struct {
	users map[string]entity.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{users: make(map[string]entity.User)}
}

func (r *UserRepository) GetUserByID(id string) (entity.User, error) {
	user, exists := r.users[id]
	if !exists {
		return entity.User{}, errors.New("usuário não encontrado")
	}
	return user, nil
}

func (r *UserRepository) CreateUser(user entity.User) {
	r.users[user.ID] = user
}
