package main

import (
	"flag"
	"fmt"
	"github.com/gfbett/gtd/tasklist"
)

var inbox = flag.String("i", "", "Adds a task in the inbox")
var removeTask = flag.Int("r", -1, "Remove a task given this index")

func main() {

	// inputs parameters declaration
	flag.Parse()

	gtd := LoadGTD()
	modified := false

	if *inbox != "" {
		fmt.Println("Adding task", *inbox)
		gtd.inbox.AddTask(tasklist.NewTask(*inbox))
		modified = true
	} else if *removeTask != -1 {
		taskIndex := *removeTask
		task := gtd.inbox.GetTask(taskIndex)
		if task != nil {
			fmt.Println("Removing task:" + task.Name)
			gtd.inbox.RemoveTask(taskIndex)
			modified = true
		}
	} else {
		fmt.Println("Current tasks:")
		for i := 0; i < gtd.inbox.Size(); i++ {
			fmt.Println(fmt.Sprint(i) + " - " + gtd.inbox.GetTask(i).Name)
		}
	}
	if modified {
		gtd.Store()
	}
}
