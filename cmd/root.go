package cmd

import (
	"github.com/spf13/cobra"
	"github.com/mitchellh/go-homedir"
	"path/filepath"
	"TaskApp/db"
)

var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI task manager",
}

func getDB() (*db.StormDB, error) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	return db.Init(dbPath)
}