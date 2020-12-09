package tasklist

import (
	"fmt"
	"testing"
)

func TestInbox_Init(t *testing.T) {
	inbox := InitTaskList()
	if cap(inbox.Tasks) != 10 {
		t.Error("Unexpected initial size")
	}
}

func TestInboxAdd(t *testing.T) {
	inbox := InitTaskList()
	name := "test name"
	inbox.AddTask(NewTask(name))
	if inbox.GetTask(0).name != name {
		t.Error("Unexpected task received")
	}
	if inbox.Size() != 1 {
		t.Error("Unexpected size received")
	}
}

func TestInboxAddMultiple(t *testing.T) {
	inbox := InitTaskList()
	taskTemplate := "Task%d"
	for i := 0; i < 12; i++ {
		taskName := fmt.Sprintf(taskTemplate, i)
		inbox.AddTask(NewTask(taskName))
	}

	if inbox.Size() != 12 {
		t.Error("Unexpected size")
	}

	for i := 0; i < 12; i++ {
		taskName := fmt.Sprintf(taskTemplate, i)
		if inbox.GetTask(i).name != taskName {
			t.Error("Unexpected task")
		}
	}
}
