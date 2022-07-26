package main

import (
	"log"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/o-ifeanyi/gophercises/go-cli-task-mgr/cmd"
	"github.com/o-ifeanyi/gophercises/go-cli-task-mgr/db"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	err := db.Init(cmd.TaskBucket, cmd.CompleteBucket, dbPath)
	if err != nil {
		log.Fatalln("Failed to init database")
	}
	err = cmd.RootCmd.Execute()
	if err != nil {
		log.Fatalln("Execution failed")
	}
}
