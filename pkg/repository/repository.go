package repository

import (
	"github.com/HACK3R911/go-todo-api/internal/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
}

type TodoList interface {
}

type TodoTask interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoTask
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
