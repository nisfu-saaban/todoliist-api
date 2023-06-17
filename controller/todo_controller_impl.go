package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nisfu-saaban/todoliist-api/app"
	"github.com/nisfu-saaban/todoliist-api/model"
	"github.com/nisfu-saaban/todoliist-api/service"
)

type TodoControllerImpl struct {
	TodoService service.TodoService
}

func NewTodoController(service service.TodoService) TodoController {
	return &TodoControllerImpl{TodoService: service}
}

func (controller *TodoControllerImpl) Create(c *gin.Context) {
	var todo model.Todo
	c.BindJSON(&todo)
	err := controller.TodoService.Create(&todo)
	if err != nil {
		app.RespondJSON(c, 404, todo)
	} else {
		app.RespondJSON(c, 200, todo)
	}
}

func (controller *TodoControllerImpl) Update(c *gin.Context) {
	var todo model.Todo
	id := c.Params.ByName("id")
	err := controller.TodoService.FindOne(&todo, id)
	if err != nil {
		app.RespondJSON(c, 404, todo)
	}
	c.BindJSON(&todo)
	err = controller.TodoService.Update(&todo, id)
	if err != nil {
		app.RespondJSON(c, 404, todo)
	} else {
		app.RespondJSON(c, 200, todo)
	}

}

func (controller *TodoControllerImpl) Delete(c *gin.Context) {
	var todo model.Todo
	id := c.Params.ByName("id")
	err := controller.TodoService.Delete(&todo, id)
	if err != nil {
		app.RespondJSON(c, 404, todo)
	} else {
		app.RespondJSON(c, 200, todo)
	}
}

func (controller *TodoControllerImpl) FindOne(c *gin.Context) {
	var todo model.Todo
	id := c.Params.ByName("id")
	err := controller.TodoService.FindOne(&todo, id)
	if err != nil {
		app.RespondJSON(c, 404, todo)
	} else {
		app.RespondJSON(c, 200, todo)
	}
}

func (controller *TodoControllerImpl) FindAll(c *gin.Context) {
	var todo []model.Todo
	err := controller.TodoService.FindAll(&todo)
	if err != nil {
		app.RespondJSON(c, 404, todo)
	} else {
		app.RespondJSON(c, 200, todo)
	}
}
