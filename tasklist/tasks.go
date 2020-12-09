package tasklist

type Task struct {
	name string
}

func NewTask(name string) *Task {
	task := new(Task)
	task.name = name
	return task
}
