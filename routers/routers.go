package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nisfu-saaban/todoliist-api/controller"
	"github.com/nisfu-saaban/todoliist-api/service"
)

func SetupRouter() *gin.Engine {
	todoService := service.NewTodoService()
	todoController := controller.NewTodoController(todoService)
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/todo", todoController.FindAll)
		v1.GET("/todo/:id", todoController.FindOne)
		v1.POST("/todo", todoController.Create)
		v1.PUT("/todo/:id", todoController.Update)
		v1.DELETE("/todo/:id", todoController.Delete)
	}
	return r
}
