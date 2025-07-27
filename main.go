package main

import (
	"fmt"
	"github.com/ikartiksh/TaskCLI/tasks"

)

func main () {
	t1 := tasks.NewTask("1", "Memory Managment")
	t2 := tasks.NewTask("2", "Usuage")


	tasks.AddTasks(t1)
	tasks.AddTasks(t2)

	allTasks := tasks.ListAllTasks()
	for _, t := range allTasks {
		fmt.Printf("Task Id : %s, Name: %s, Status: %s\n", t.ID, t.Name, t.Status)
	}
}


