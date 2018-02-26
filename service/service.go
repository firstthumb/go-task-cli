package service

import "TaskApp/entity"

type TaskService interface {
	Create(task *entity.Task) (int, error)
	GetAll() ([]*entity.Task, error)
}
