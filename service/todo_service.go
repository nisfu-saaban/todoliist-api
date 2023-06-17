package service

import "github.com/nisfu-saaban/todoliist-api/model"

type TodoService interface {
	Create(model *model.Todo) error
	Update(model *model.Todo, id string) error
	Delete(model *model.Todo, id string) error
	FindOne(model *model.Todo, id string) error
	FindAll(model *[]model.Todo) error
}
