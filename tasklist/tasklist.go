package tasklist

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type TaskList struct {
	tasks []*Task
}

func (list *TaskList) Size() int {
	return len(list.tasks)
}

func (list *TaskList) AddTask(task *Task) {
	list.tasks = append(list.tasks, task)
}

func (list *TaskList) GetTask(i int) *Task {
	if i < 0 || i >= list.Size() {
		return nil
	}
	return list.tasks[i]
}

func (list *TaskList) RemoveTask(i int) bool {
	if i < 0 || i >= list.Size() {
		return false
	}
	list.tasks = append(list.tasks[:i], list.tasks[i+1:]...)
	return true
}

func InitTaskList() *TaskList {
	list := new(TaskList)
	list.tasks = make([]*Task, 0, 10)
	return list
}

func (list *TaskList) ToStorableString() string {
	var b strings.Builder
	size := list.Size()
	fmt.Fprint(&b, size)
	for i := 0; i < size; i++ {
		task := list.tasks[i]
		if !task.Completed() {
			b.WriteString("\n")
			fmt.Fprint(&b, list.tasks[i].ToStorableString())
		}
	}
	for i := 0; i < size; i++ {
		task := list.tasks[i]
		if task.Completed() {
			b.WriteString("\n")
			fmt.Fprint(&b, list.tasks[i].ToStorableString())
		}
	}
	return b.String()
}

func (list *TaskList) LoadFromStorableString(data string) bool {
	parts := strings.Split(data, "\n")
	size, err := strconv.Atoi(parts[0])
	if err != nil || size < 0 || len(parts) < size+1 {
		log.Print("Invalid task count")
		return false
	}
	list.tasks = make([]*Task, 0, size)
	for i := 0; i < size; i++ {
		task := LoadTask(parts[i+1])
		if task == nil {
			log.Print("Unable to load task from: " + parts[i+1])
			return false
		}
		list.AddTask(task)
	}
	return true
}

func (list *TaskList) FileName() string {
	return "inbox.gtd"
}
