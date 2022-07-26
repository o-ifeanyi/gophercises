# CLI Task Manager

[Source](https://courses.calhoun.io/lessons/les_goph_35)

## Exercise details

A CLI tool that can be used to manage your TODOs in the terminal. The basic usage of the tool is going to look roughly like this:

$ task

Task is a CLI task manager.

Usage:

- task [command]

Available Commands:

- add - Adds a new task to your list
- do - Marks a task as complete
- list - List all of your incomplete tasks
- completed - List all of your completed tasks
- rm - Removes a task from your list
- reset - Resets all task list and database

Use "task [command] --help" for more information about a command.

$ task add review talk proposal
- Added "review talk proposal" to your list.

$ task add clean dishes
- Added "clean dishes" to your list.

$ task add build a cli in GO
- Added "build a cli in GO" to your list.

$ task list
1. review talk proposal
2. clean dishes
3. build a cli in GO

$ task do 1
- Marked "review talk proposal" as complete.

$ task list
1. clean dishes
2. build a cli in GO

$ task completed
1. "review talk proposal" completed.

$ task rm 1
- Removed "clean dishes" from your list.

$ task list
1. build a cli in GO

$ task completed
1. "review talk proposal" completed.

$ task reset
- Reset successful!

$ task list
- You currently have no task.

$ task completed
- You currently have no completed task.
