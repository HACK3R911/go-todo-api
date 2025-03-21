package models

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoTask struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListsTask struct {
	Id     int
	ListId int
	TaskId int
}
