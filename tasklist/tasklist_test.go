package tasklist

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	inbox := InitTaskList()
	if cap(inbox.tasks) != 10 {
		t.Error("Unexpected initial size")
	}
}

func TestAdd(t *testing.T) {
	inbox := InitTaskList()
	name := "test name"
	inbox.AddTask(NewTask(name))
	if inbox.GetTask(0).Name != name {
		t.Error("Unexpected task received")
	}
	if inbox.Size() != 1 {
		t.Error("Unexpected size received")
	}
}

func TestAddMultiple(t *testing.T) {
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
		if inbox.GetTask(i).Name != taskName {
			t.Error("Unexpected task")
		}
	}
}

func TestRemove(t *testing.T) {
	inbox := InitTaskList()
	taskTemplate := "Task%d"
	for i := 0; i < 10; i++ {
		taskName := fmt.Sprintf(taskTemplate, i)
		inbox.AddTask(NewTask(taskName))
	}
	if inbox.RemoveTask(-1) != false {
		t.Error("Remove with negative index should have failed")
	}
	if inbox.RemoveTask(10) != false {
		t.Error("Remove with index too large should have failed")
	}
	if inbox.RemoveTask(5) != true {
		t.Error("Remove with correct index should succeed")
	}
	var expected = [...]string{"Task0", "Task1", "Task2", "Task3", "Task4", "Task6", "Task7", "Task8", "Task9"}
	for i := 0; i < inbox.Size(); i++ {
		taskName := inbox.GetTask(i).Name
		if taskName != expected[i] {
			t.Error("Unexpected name, expected: " + expected[i] + "found: " + taskName)
		}
	}
}

func TestToStorableString(t *testing.T) {
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

func TestLoadInvalidSize(t *testing.T) {
	storable := "4\nTask1\nTask2\nTask3"
	taskList := InitTaskList()
	taskList.LoadFromStorableString(storable)
	size := taskList.Size()
	if size != 0 {
		t.Error("Unexpected size: " + fmt.Sprint(size))
	}
}

func TestLoadSizeNotANumber(t *testing.T) {
	storable := "lala\nTask1\nTask2\nTask3"
	taskList := InitTaskList()
	taskList.LoadFromStorableString(storable)
	size := taskList.Size()
	if size != 0 {
		t.Error("Unexpected size: " + fmt.Sprint(size))
	}
}

func TestLoadNegativeSize(t *testing.T) {
	storable := "-3\nTask1\nTask2\nTask3"
	taskList := InitTaskList()
	taskList.LoadFromStorableString(storable)
	size := taskList.Size()
	if size != 0 {
		t.Error("Unexpected size: " + fmt.Sprint(size))
	}
}

func TestLoadInvalidData(t *testing.T) {
	storable := ""
	taskList := InitTaskList()
	taskList.LoadFromStorableString(storable)
	size := taskList.Size()
	if size != 0 {
		t.Error("Unexpected size: " + fmt.Sprint(size))
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
		name := taskList.GetTask(i).Name
		if name != "Task"+fmt.Sprint(i+1) {
			t.Error("Unexpected task name: " + name)
		}
	}
}
