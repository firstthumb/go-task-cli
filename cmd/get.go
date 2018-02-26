package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets the task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		taskService := getTaskService()
		taskId, _ := strconv.Atoi(args[0])
		task, err := taskService.Get(taskId)
		if err != nil {
			fmt.Println("Something went wrong", err)
			return
		}
		fmt.Printf("Task \"%s\" from your task list with ID : %d\n", task.Value, taskId)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
