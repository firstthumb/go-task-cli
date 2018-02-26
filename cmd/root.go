package cmd

import (
	"github.com/spf13/cobra"
	"github.com/mitchellh/go-homedir"
	"path/filepath"
	"TaskApp/repository"
	"TaskApp/service"
)

var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI task manager",
}

func getTaskService() service.TaskService {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	taskRepo := repository.NewStormDBRepository(dbPath)
	return service.NewService(taskRepo)
}