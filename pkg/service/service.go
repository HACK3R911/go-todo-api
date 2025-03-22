package service

import "github.com/HACK3R911/go-todo-api/pkg/repository"

type Authorization interface {
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
	return &Service{}
}
