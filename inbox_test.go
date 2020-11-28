package main

import (
	"testing"
	"fmt"
)

func TestInbox_Init(t *testing.T) {
	inbox := InitInbox()
	if cap(inbox.Tasks) != 10 {
		t.Error("Unexpected initial size")
	}
}

func TestInboxAdd(t *testing.T) {
	inbox := InitInbox()
	task := "test task"
	inbox.AddTask(task)
	if inbox.GetTask(0) != task {
		t.Error("Unexpected task received")
	}
	if inbox.Size() != 1 {
		t.Error("Unexpected size received")
	}
}

func TestInboxAddMultiple(t *testing.T) {
	inbox := InitInbox()
	taskTemplate := "Task%d"
	for i :=0; i< 12; i++ {
		taskName := fmt.Sprintf(taskTemplate, i)
		inbox.AddTask(taskName)
	}

	if inbox.Size() != 12 {
		t.Error("Unexpected size")
	}

	for i :=0; i< 12; i++ {
		taskName := fmt.Sprintf(taskTemplate, i)
		if inbox.GetTask(i) != taskName {
			t.Error("Unexpected task")
		}
	}
}

