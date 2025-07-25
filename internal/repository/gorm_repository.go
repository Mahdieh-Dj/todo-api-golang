package repository

import (
	"todo-app/internal/model"

	"gorm.io/gorm"
)

type GormRepository struct {
	DB *gorm.DB
}

func(r *GormRepository)Create(todo *model.Todo) error{
	return r.DB.Create(todo).Error
}

func (r *GormRepository)GetAll()([]model.Todo,error){
	var todos []model.Todo
	err:=r.DB.Find(&todos).Error

	return todos,err
}

func (r *GormRepository)GetByID(id string)(*model.Todo,error){
	var todo model.Todo
	err:=r.DB.First(&todo,id).Error

	return &todo,err
}

func (r *GormRepository)Update(todo *model.Todo)error{
	return  r.DB.Save(todo).Error
}

func (r *GormRepository) Delete(id string) error {
	return r.DB.Delete(&model.Todo{}, id).Error
}