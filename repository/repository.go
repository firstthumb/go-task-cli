package repository

import "TaskApp/entity"

type Repository interface {
	//Find(id int) (*entity.Task, error)
	FindAll() ([]*entity.Task, error)
	Store(task *entity.Task) (int, error)
}
