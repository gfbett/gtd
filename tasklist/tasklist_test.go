package tasklist

import (
	"fmt"
	"testing"
)

func TestInbox_Init(t *testing.T) {
	inbox := InitTaskList()
	if cap(inbox.tasks) != 10 {
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

func TestTaskListToStorableString(t *testing.T) {
	taskList := InitTaskList()
	storable := taskList.ToStorableString()
	if storable != "0" {
		t.Error("Unexpected storable string")
	}
	taskList.AddTask(NewTask("Test task"))
	taskList.AddTask(NewTask("Another task"))
	storable = taskList.ToStorableString()
	if storable != "2\nTest task\nAnother task" {
		t.Error("Unexpected storable string:" + storable)
	}
}

func TestTaskListLoadFromStorableString(t *testing.T) {
	storable := "3\nTask1\nTask2\nTask3"
	taskList := InitTaskList()
	taskList.LoadFromStorableString(storable)
	size := taskList.Size()
	if size != 3 {
		t.Error("Unexpected size: " + fmt.Sprint(size))
	}
	for i := 0; i < 3; i++ {
		name := taskList.GetTask(i).name
		if name != "Task"+fmt.Sprint(i+1) {
			t.Error("Unexpected task name: " + name)
		}
	}
}
