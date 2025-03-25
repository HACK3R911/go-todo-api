package repository

import (
	"fmt"
	"github.com/HACK3R911/go-todo-api/internal/models"
	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userId int, list models.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *TodoListPostgres) GetAll(userId int) ([]models.TodoList, error) {
	var lists []models.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1",
		todoListsTable, usersListsTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}
