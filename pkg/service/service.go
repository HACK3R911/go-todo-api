package service

import (
	"github.com/HACK3R911/go-todo-api/internal/models"
	"github.com/HACK3R911/go-todo-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
}

type TodoList interface {
}

type TodoTask interface {
}

type Service struct {
	Authorization
	TodoList
	TodoTask
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
