package tasklist

import (
	"strings"
)

type Task struct {
	name string
}

func NewTask(name string) *Task {
	task := new(Task)
	task.name = strings.TrimSpace(name)
	return task
}

func (task *Task) ToStorableString() string {
	return task.name
}

func (task *Task) Name() string {
	return task.name
}

func (task *Task) SetName(name string) {
	task.name = strings.TrimSpace(name)
}
