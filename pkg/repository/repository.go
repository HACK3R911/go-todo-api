package repository

import (
	"github.com/HACK3R911/go-todo-api/internal/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type TodoList interface {
	Create(userId int, list models.TodoList) (int, error)
	GetAll(userId int) ([]models.TodoList, error)
	GetById(userId, listId int) (models.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input models.UpdateListInput) error
}

type TodoTask interface {
	Create(listId int, task models.TodoTask) (int, error)
	GetAll(userId, listId int) ([]models.TodoTask, error)
	GetById(userId, taskId int) (models.TodoTask, error)
	Delete(userId, listId int) error
	//Update(userId, taskId int, input models.UpdateTaskInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoTask
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoTask:      NewTodoTaskPostgres(db),
	}
}
