package service

import (
	"github.com/HACK3R911/go-todo-api/internal/models"
	"github.com/HACK3R911/go-todo-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list models.TodoList) (int, error)
	GetAll(userId int) ([]models.TodoList, error)
	GetById(userId, listId int) (models.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input models.UpdateListInput) error
}

type TodoTask interface {
	Create(userId, listId int, task models.TodoTask) (int, error)
	GetAll(userId, listId int) ([]models.TodoTask, error)
	GetById(userId, taskId int) (models.TodoTask, error)
	Delete(userId, taskId int) error
	Update(userId, taskId int, input models.UpdateTaskInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoTask
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoTask:      NewTodoTaskService(repos.TodoTask, repos.TodoList),
	}
}
