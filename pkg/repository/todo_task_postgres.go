package repository

import (
	"fmt"
	"github.com/HACK3R911/go-todo-api/internal/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

type TodoTaskPostgres struct {
	db *sqlx.DB
}

func NewTodoTaskPostgres(db *sqlx.DB) *TodoTaskPostgres {
	return &TodoTaskPostgres{db: db}
}

func (r *TodoTaskPostgres) Create(listId int, task models.TodoTask) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var taskId int
	createTaskQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoTasksTable)

	row := tx.QueryRow(createTaskQuery, task.Title, task.Description)
	err = row.Scan(&taskId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListTaskQuery := fmt.Sprintf("INSERT INTO %s (list_id, task_id) VALUES ($1, $2)", listsTasksTable)
	_, err = tx.Exec(createListTaskQuery, listId, taskId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return taskId, tx.Commit()
}

func (r *TodoTaskPostgres) GetAll(userId, listId int) ([]models.TodoTask, error) {
	var tasks []models.TodoTask
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li ON li.task_id = ti.id 
    								INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2`,
		todoTasksTable, listsTasksTable, usersListsTable)
	if err := r.db.Select(&tasks, query, userId, listId); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TodoTaskPostgres) GetById(userId, taskId int) (models.TodoTask, error) {
	var task models.TodoTask
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li ON li.task_id = ti.id 
    								INNER JOIN %s ul on ul.list_id = li.list_id WHERE ti.id = $1 AND ul.user_id = $2`,
		todoTasksTable, listsTasksTable, usersListsTable)
	if err := r.db.Get(&task, query, userId, taskId); err != nil {
		return task, err
	}

	return task, nil
}

func (r *TodoTaskPostgres) Delete(userId, taskId int) error {
	query := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul 
       								WHERE ti.id = li.task_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2`,
		todoTasksTable, listsTasksTable, usersListsTable)

	_, err := r.db.Exec(query, userId, taskId)
	return err
}

func (r *TodoTaskPostgres) Update(userId, taskId int, input models.UpdateTaskInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ti SET %s  FROM %s li, %s ul WHERE ti.id = li.task_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d`,
		todoTasksTable, setQuery, listsTasksTable, usersListsTable, argId, argId+1)

	args = append(args, userId, taskId)

	_, err := r.db.Exec(query, args...)
	return err
}
