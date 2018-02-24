package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		database, _ := getDB()
		task := strings.Join(args, " ")
		taskId, err := database.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong", err)
			return
		}
		fmt.Printf("Added \"%s\" to your task list with ID : %d\n", task, taskId)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
