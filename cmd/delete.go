package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes the task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		taskService := getTaskService()
		taskId, _ := strconv.Atoi(args[0])
		task, err := taskService.Delete(taskId)
		if err != nil {
			fmt.Println("Something went wrong", err)
			return
		}
		fmt.Printf("Deleted \"%s\" from your task list with ID : %d\n", task.Value, taskId)
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
