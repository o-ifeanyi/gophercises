package cmd

import "github.com/spf13/cobra"

var TaskBucket = []byte("tasks")
var CompleteBucket = []byte("complete")
var RootCmd = &cobra.Command{Use: "task",
	Short: "Task is a CLI task manager"}
