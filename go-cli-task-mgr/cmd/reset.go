package cmd

import (
	"fmt"
	"log"

	"github.com/o-ifeanyi/gophercises/go-cli-task-mgr/db"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Resets all task list and database",
	Run: func(cmd *cobra.Command, args []string) {
		err := db.ResetTask(TaskBucket)
		if err != nil {
			log.Fatalln("Something went wrong:", err)
		}
		err = db.ResetTask(CompleteBucket)
		if err != nil {
			log.Fatalln("Something went wrong:", err)
		}
		fmt.Printf("\nReset successful!\n\n")
	},
}

func init() {
	RootCmd.AddCommand(resetCmd)
}
