package tasklist

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Task struct {
	name      string
	completed bool
}

func NewTask(name string) *Task {
	task := new(Task)
	task.name = strings.TrimSpace(name)
	task.completed = false
	return task
}

func LoadTask(storedString string) *Task {
	task := new(Task)
	parts := strings.Split(storedString, "|")
	if len(parts) != 2 {
		log.Print("Not enough parts for loading task")
		return nil
	}
	task.name = parts[0]
	completed, err := strconv.ParseBool(parts[1])
	if err != nil {
		log.Print("Completed not a valid bool")
		return nil
	}
	task.completed = completed

	return task
}

func (task *Task) ToStorableString() string {
	return fmt.Sprintf("%s|%t", task.name, task.completed)
}

func (task *Task) Completed() bool {
	return task.completed
}

func (task *Task) SetCompleted(completed bool) {
	task.completed = completed
}

func (task *Task) Name() string {
	return task.name
}

func (task *Task) SetName(name string) {
	task.name = strings.TrimSpace(name)
}
