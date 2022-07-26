package cmd

import (
	"fmt"
	"log"

	"github.com/o-ifeanyi/gophercises/go-cli-task-mgr/db"
	"github.com/spf13/cobra"
)

var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "List all of your completed tasks",
	Run: func(cmd *cobra.Command, args []string) {
		allTask, err := db.AllTasks(CompleteBucket)
		if err != nil {
			log.Fatalln("Something went wrong:", err)
		}
		if len(allTask) < 1 {
			fmt.Printf("\nYou currently have no completed task.\n\n")
		}
		for _, task := range allTask {
			fmt.Printf("\n%d. \"%s\" completed.\n\n", task.Key, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(completedCmd)
}
