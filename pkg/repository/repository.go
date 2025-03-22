package repository

type Authorization interface {
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

func NewRepository() *Repository {
	return &Repository{}
}
