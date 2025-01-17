package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
	"TaskApp/entity"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		taskService := getTaskService()
		task := strings.Join(args, " ")
		taskId, err := taskService.Create(&entity.Task{Value: task})
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
