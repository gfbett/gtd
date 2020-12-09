package tasklist

type Task struct {
	Name string
}

func NewTask(name string) *Task {
	task := new(Task)
	task.Name = name
	return task
}

func (task *Task) ToStorableString() string {
	return task.Name
}
