package models

import (
	"Bubble/dao"
)

// Todo 1.定义模型
type Todo struct {
	ID 		int 	`json:"id"`
	Title 	string 	`json:"title"`
	Status 	bool  	`json:"status"`
}

/*
  Todo这个Model增删改查操作都放在这里
*/

//CreateATodo 创建todo
func CreateATodo(todo *Todo)(err error){
	 err = dao.DB.Create(&todo).Error
	return
}

func GetALlTodo()(todoList []*Todo,err error){
	if err =dao.DB.Find(&todoList).Error;err!=nil{
		return nil ,err
	}
	return
}

func GetATodo(id string)(todo *Todo,err error){
	todo=new(Todo)
	if err=dao.DB.Where("id=?",id).First(todo).Error;err!=nil{
		return nil,err
	}
	return
}

func UpdateATodo(todo *Todo)(err error){
	 err =dao.DB.Save(todo).Error
	return
}

func DeleteAtodo(id string)(err error){
	err=dao.DB.Where("id=?",id).Delete(&Todo{}).Error
	return
}
