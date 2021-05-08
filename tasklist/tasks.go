package tasklist

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	name          string
	completed     bool
	createdDate   time.Time
	completedDate time.Time
}

func NewTask(name string) *Task {
	task := new(Task)
	task.name = strings.TrimSpace(name)
	task.completed = false
	task.createdDate = time.Now()
	return task
}

func LoadTask(storedString string) *Task {
	task := new(Task)
	parts := strings.Split(storedString, "|")
	if len(parts) < 2 {
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
	if len(parts) >= 3 {
		var date time.Time
		err = date.UnmarshalText([]byte(parts[2]))
		if err != nil {
			log.Print("Cannot convert created date")
			return nil
		}
		task.createdDate = date
	}
	if len(parts) >= 4 {
		var date time.Time
		err = date.UnmarshalText([]byte(parts[3]))
		if err != nil {
			log.Print("Cannot convert completed date")
			return nil
		}
		task.completedDate = date
	}

	return task
}

func (task *Task) ToStorableString() string {
	createdDate, _ := task.createdDate.MarshalText()
	completedDate, _ := task.completedDate.MarshalText()
	return fmt.Sprintf("%s|%t|%s|%s", task.name, task.completed, createdDate, completedDate)
}

func (task *Task) Completed() bool {
	return task.completed
}

func (task *Task) CreatedDate() time.Time {
	return task.createdDate
}

func (task *Task) CompletedDate() time.Time {
	return task.completedDate
}

func (task *Task) SetCompleted(completed bool) {
	task.completed = completed
	if completed {
		task.completedDate = time.Now()
	} else {
		var zeroTime time.Time
		task.completedDate = zeroTime
	}
}

func (task *Task) Name() string {
	return task.name
}

func (task *Task) SetName(name string) {
	task.name = strings.TrimSpace(name)
}
