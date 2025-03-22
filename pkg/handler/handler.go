package handler

import (
	"github.com/HACK3R911/go-todo-api/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			tasks := lists.Group(":id/tasks")
			{
				tasks.POST("/", h.createTask)
				tasks.GET("/", h.getAllTasks)
				tasks.GET("/:task_id", h.getTaskById)
				tasks.PUT("/:task_id", h.updateTask)
				tasks.DELETE("/:task_id", h.deleteTask)
			}
		}
	}
	return router
}
