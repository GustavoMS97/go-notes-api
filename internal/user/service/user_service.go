package service

import (
	"github.com/GustavoMS97/go-notes-api/internal/user/entity"
	"github.com/GustavoMS97/go-notes-api/internal/user/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(id string) (entity.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserService) CreateUser(user entity.User) {
	s.repo.CreateUser(user)
}
