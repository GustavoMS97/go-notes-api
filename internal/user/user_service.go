package user

import (
	"errors"
	"log"

	"github.com/GustavoMS97/go-notes-api/internal/auth"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(name, email, plainPassword string) (User, error) {
	log.Println("[UserService] Creating user:", email)

	existing, _ := s.repo.FindByEmail(email)
	if existing != nil {
		log.Println("[UserService] Email already exists:", email)
		return User{}, errors.New("email already registered")
	}

	hashedPassword, err := auth.HashPassword(plainPassword)
	if err != nil {
		log.Println("[UserService] Failed to hash password:", err)
		return User{}, err
	}

	user := User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	createdUser, err := s.repo.Create(user)
	if err != nil {
		log.Println("[UserService] Failed to create user:", err)
		return User{}, err
	}

	log.Println("[UserService] User created successfully:", createdUser.ObjectID)
	return createdUser, nil
}
