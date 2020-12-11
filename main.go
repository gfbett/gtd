package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/gfbett/gtd/tasklist"
	"os"
)

var inbox = flag.String("i", "", "Adds a task in the inbox")
var removeTask = flag.Int("r", -1, "Remove a task given this index")
var editTask = flag.Int("e", -1, "Edit a task given this index")

func main() {

	// inputs parameters declaration
	flag.Parse()

	gtd := LoadGTD()
	if gtd == nil {
		fmt.Println("Unable to read stored tasks")
		os.Exit(1)
	}
	modified := false

	if *inbox != "" {
		fmt.Println("Adding task", *inbox)
		gtd.inbox.AddTask(tasklist.NewTask(*inbox))
		modified = true
	} else if *removeTask != -1 {
		taskIndex := *removeTask
		modified = RemoveTask(gtd.inbox, taskIndex)
	} else if *editTask != -1 {
		taskIndex := *editTask
		modified = EditTask(gtd.inbox, taskIndex)
	} else {
		ListTasks(gtd.inbox)
	}
	if modified {
		gtd.Store()
	}

}

func RemoveTask(tasklist *tasklist.TaskList, index int) bool {
	task := tasklist.GetTask(index)
	if task != nil {
		fmt.Println("Removing task:" + task.Name())
		tasklist.RemoveTask(index)
		return true
	}
	return false
}

func EditTask(tasklist *tasklist.TaskList, index int) bool {
	reader := bufio.NewReader(os.Stdin)
	task := tasklist.GetTask(index)
	if task != nil {
		fmt.Println("Current task description:", task.Name())
		fmt.Println("Enter new description")
		name, _ := reader.ReadString('\n')
		fmt.Println("New description:", name)
		task.SetName(name)
		return true
	}
	return false
}

func ListTasks(list *tasklist.TaskList) {
	fmt.Println("Current tasks:")
	for i := 0; i < list.Size(); i++ {
		fmt.Println(fmt.Sprintf("%3d - %s", i, list.GetTask(i).Name()))
	}
}
