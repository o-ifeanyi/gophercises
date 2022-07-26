# CLI Task Manager

[Source](https://courses.calhoun.io/lessons/les_goph_35)

## Exercise details

A CLI tool that can be used to manage your TODOs in the terminal. The basic usage of the tool is going to look roughly like this:

$ go-cli-task-mgr

Task is a CLI task manager.

Usage:

- go-cli-task-mgr [command]

Available Commands:

- add - Adds a new task to your list
- do - Marks a task as complete
- list - List all of your incomplete tasks
- completed - List all of your completed tasks
- rm - Removes a task from your list
- reset - Resets all task list and database

Use "go-cli-task-mgr [command] --help" for more information about a command.

$ go-cli-task-mgr add review talk proposal
- Added "review talk proposal" to your list.

$ go-cli-task-mgr add clean dishes
- Added "clean dishes" to your list.

$ go-cli-task-mgr add build a cli in GO
- Added "build a cli in GO" to your list.

$ go-cli-task-mgr list
1. review talk proposal
2. clean dishes
3. build a cli in GO

$ go-cli-task-mgr do 1
- Marked "review talk proposal" as complete.

$ go-cli-task-mgr list
1. clean dishes
2. build a cli in GO

$ go-cli-task-mgr completed
1. "review talk proposal" completed.

$ go-cli-task-mgr rm 1
- Removed "clean dishes" from your list.

$ go-cli-task-mgr list
1. build a cli in GO

$ go-cli-task-mgr completed
1. "review talk proposal" completed.

$ go-cli-task-mgr reset
- Reset successful!

$ go-cli-task-mgr list
- You currently have no task.

$ go-cli-task-mgr completed
- You currently have no completed task.
