package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/gfbett/gtd/tasklist"
	"os"
)

var inbox = flag.Bool("i", false, "Adds a task in the inbox")
var completeTask = flag.Int("c", -1, "Complete a task given this index")
var editTask = flag.Int("e", -1, "Edit a task given this index")

func main() {

	// inputs parameters declaration
	flag.Parse()

	gtd := LoadGTD()
	modified := false

	if *inbox != false {
		modified = AddTask(gtd.inbox)
	} else if *completeTask != -1 {
		taskIndex := *completeTask
		modified = CompleteTask(gtd.inbox, taskIndex)
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

func AddTask(inbox *tasklist.TaskList) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter new task description: ")
	name, _ := reader.ReadString('\n')
	if len(name) > 0 {
		inbox.AddTask(tasklist.NewTask(name))
		fmt.Println("New Task:", name)
		return true
	} else {
		return false
	}
}

func CompleteTask(tasklist *tasklist.TaskList, index int) bool {
	task := tasklist.GetTask(index)
	if task != nil {
		fmt.Println("Complete task:" + task.Name())
		task.SetCompleted(true)
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
		task := list.GetTask(i)
		if !task.Completed() {
			fmt.Println(fmt.Sprintf("%3d - %s", i, task.Name()))
		}

	}
}
