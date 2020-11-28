package main

type Inbox struct {
	Tasks []string
}

func (inbox *Inbox) Size() int {
	return len(inbox.Tasks)
}

func (inbox *Inbox) AddTask(task string) {
	inbox.Tasks = append(inbox.Tasks, task)
}

func (inbox *Inbox) GetTask(i int) string {
	return inbox.Tasks[i]
}

func InitInbox() *Inbox {
	inbox := new(Inbox)
	inbox.Tasks = make([]string, 0, 10)
	return inbox
}
