package service

import (
	"github.com/HACK3R911/go-todo-api/internal/models"
	"github.com/HACK3R911/go-todo-api/pkg/repository"
)

type TodoTaskService struct {
	repo     repository.TodoTask
	listRepo repository.TodoList
}

func NewTodoTaskService(repo repository.TodoTask, listRepo repository.TodoList) *TodoTaskService {
	return &TodoTaskService{repo: repo, listRepo: listRepo}
}

func (s *TodoTaskService) Create(userId, listId int, task models.TodoTask) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		// list not exists
		return 0, err
	}

	return s.repo.Create(listId, task)
}

func (s *TodoTaskService) GetAll(userId, listId int) ([]models.TodoTask, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *TodoTaskService) GetById(userId, taskId int) (models.TodoTask, error) {
	return s.repo.GetById(userId, taskId)
}

func (s *TodoTaskService) Delete(userId, taskId int) error {
	return s.repo.Delete(userId, taskId)
}

func (s *TodoTaskService) Update(userId, taskId int, input models.UpdateTaskInput) error {
	return s.repo.Update(userId, taskId, input)
}
