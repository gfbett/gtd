package tasklist

import (
	"fmt"
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
	return list.tasks[i]
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
		b.WriteString("\n")
		fmt.Fprint(&b, list.tasks[i].ToStorableString())
	}
	return b.String()

}

func (list *TaskList) LoadFromStorableString(data string) {
	parts := strings.Split(data, "\n")
	if len(parts) < 1 {
		return
	}
	size, err := strconv.Atoi(parts[0])
	if err != nil || len(parts) != size+1 {
		return
	}
	if size < 0 {
		size = 10
	}
	list.tasks = make([]*Task, 0, size)
	for i := 0; i < size; i++ {
		list.AddTask(NewTask(parts[i+1]))
	}
}

func (list *TaskList) FileName() string {
	return "inbox.json"
}
