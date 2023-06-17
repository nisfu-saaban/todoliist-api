package service

import (
	"github.com/nisfu-saaban/todoliist-api/app"
	"github.com/nisfu-saaban/todoliist-api/model"
)

type TodoServiceImpl struct {
}

func NewTodoService() TodoService {
	return &TodoServiceImpl{}
}

func (service *TodoServiceImpl) Create(todo *model.Todo) error {
	db := app.GetDB()

	err := db.Create(todo).Error
	if err != nil {
		return err
	}

	return nil
}

func (service *TodoServiceImpl) Update(todo *model.Todo, id string) error {
	db := app.GetDB()
	err := db.Where("id = ?", id).Save(todo).Error
	if err != nil {
		return err
	}
	return nil
}

func (service *TodoServiceImpl) Delete(todo *model.Todo, id string) error {
	db := app.GetDB()
	err := db.Where("id = ?", id).Delete(todo).Error
	if err != nil {
		return err
	}
	return nil
}

func (service *TodoServiceImpl) FindOne(todo *model.Todo, id string) error {
	db := app.GetDB()
	err := db.Where("id = ?", id).First(todo).Error
	if err != nil {
		return err
	}
	return nil
}

func (service *TodoServiceImpl) FindAll(todo *[]model.Todo) error {
	db := app.GetDB()
	err := db.Find(todo).Error
	if err != nil {
		return err
	}
	return nil
}
