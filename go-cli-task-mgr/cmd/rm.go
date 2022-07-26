package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/o-ifeanyi/gophercises/go-cli-task-mgr/db"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Removes a task from your list",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, v := range args {
			id, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalln("Something went wrong:", err)
			} else {
				ids = append(ids, id)
			}
		}

		allTask, err := db.AllTasks(TaskBucket)
		if err != nil {
			log.Fatalln("Something went wrong:", err)
		}

		for _, id := range ids {
			if id <= 0 || id > len(allTask) {
				fmt.Printf("\ninvalid task number\n\n")
				continue
			}
			task := allTask[id-1]
			err := db.DeleteTask(TaskBucket, task.Key)
			if err != nil {
				log.Fatalln("Something went wrong:", err)
			}
			fmt.Printf("\nRemoved \"%s\" from your list.\n\n", task.Value)
		}

	},
}

func init() {
	RootCmd.AddCommand(rmCmd)
}
