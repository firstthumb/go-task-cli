package service

import (
	"TaskApp/entity"
	"TaskApp/repository"
)

type TaskServiceImpl struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) TaskService {
	return &TaskServiceImpl{repo:repo}
}

func (srv *TaskServiceImpl) Create(task *entity.Task) (int, error) {
	return srv.repo.Store(task)
}

func (srv *TaskServiceImpl) GetAll() ([]*entity.Task, error) {
	return srv.repo.FindAll()
}
