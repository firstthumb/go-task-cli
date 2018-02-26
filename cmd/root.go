package cmd

import (
	"github.com/spf13/cobra"
	"github.com/mitchellh/go-homedir"
	"path/filepath"
	"TaskApp/repository"
)

var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI task manager",
}

func getDB() (*repository.StormDB, error) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	return repository.Init(dbPath)
}