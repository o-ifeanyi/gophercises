package cmd

import (
	"fmt"
	"log"

	"github.com/o-ifeanyi/gophercises/go-cli-task-mgr/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Run: func(cmd *cobra.Command, args []string) {
		allTask, err := db.AllTasks(TaskBucket)
		if err != nil {
			log.Fatalln("Something went wrong:", err)
		}
		if len(allTask) < 1 {
			fmt.Printf("\nYou currently have no task.\n\n")
		}
		for i, task := range allTask {
			fmt.Printf("\n%d. %s.\n\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
