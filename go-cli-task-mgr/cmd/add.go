package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/o-ifeanyi/gophercises/go-cli-task-mgr/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(TaskBucket, task)
		if err != nil {
			log.Fatalln("Something went wrong:", err)
		}
		fmt.Printf("\nAdded \"%s\" to your list.\n\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
