package tasklist

import (
	"testing"
)

func TestNewTask(t *testing.T) {
	expected := "Test task"
	task := NewTask(expected)
	if task == nil {
		t.Error("Task shoudn't be nil")
	}
	if task.Name() != expected {
		t.Error("Task doesn't have expected name")
	}
	if task.Completed() != false {
		t.Error("Task completed not false")
	}
}

func TestLoadTask(t *testing.T) {
	stored := "Test task|true"
	task := LoadTask(stored)
	if task == nil {
		t.Error("Task shoudn't be nil")
	}
	if task.Name() != "Test task" {
		t.Error("Task doesn't have expected name")
	}
	if task.Completed() != true {
		t.Error("Task completed not false")
	}
}

func TestStore(t *testing.T) {
	expected := "test task|false"
	task := NewTask("test task")
	storable := task.ToStorableString()
	if expected != storable {
		t.Error("Storable string not in expected format")
	}
}
