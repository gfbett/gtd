package tasklist

import (
	"fmt"
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {
	expected := "Test task"
	before := time.Now()
	task := NewTask(expected)
	after := time.Now()
	if task == nil {
		t.Error("Task shoudn't be nil")
	}
	if task.Name() != expected {
		t.Error("Task doesn't have expected name")
	}
	if task.Completed() != false {
		t.Error("Task completed not false")
	}
	if !(task.CreatedDate().After(before) && task.CreatedDate().Before(after)) {
		t.Error("Task creation date not the expected")
	}
}

func TestLoadTask(t *testing.T) {
	stored := "Test task|true|2021-01-01T01:01:01Z|2022-01-01T01:01:01Z"
	createdDate := time.Date(2021, 01, 01, 01, 01, 01, 0, time.UTC)
	completedDate := time.Date(2022, 01, 01, 01, 01, 01, 0, time.UTC)
	task := LoadTask(stored)
	if task == nil {
		t.Error("Task shouldn't be nil")
	}
	if task.Name() != "Test task" {
		t.Error("Task doesn't have expected name")
	}
	if task.Completed() != true {
		t.Error("Task completed not false")
	}
	if task.CreatedDate() != createdDate {
		t.Error("Task created date not the expected")
	}
	if task.CompletedDate() != completedDate {
		t.Error("Task completed date not the expected")
	}
}

func TestLoadTaskOld(t *testing.T) {
	stored := "Test task|true"
	task := LoadTask(stored)
	if task == nil {
		t.Error("Task shouldn't be nil")
	}
	if task.Name() != "Test task" {
		t.Error("Task doesn't have expected name")
	}
	if task.Completed() != true {
		t.Error("Task completed not false")
	}
	if !task.CreatedDate().IsZero() {
		t.Error("Task created date not the expected")
	}
	if !task.CompletedDate().IsZero() {
		t.Error("Task completed date not the expected")
	}
}

func TestStore(t *testing.T) {
	expected := "test task|false|2021-01-01T01:01:01Z|0001-01-01T00:00:00Z"
	task := NewTask("test task")
	task.createdDate = time.Date(2021, 01, 01, 01, 01, 01, 0, time.UTC)
	storable := task.ToStorableString()
	if expected != storable {
		t.Error(fmt.Sprintf("Storable string not in expected format: \nExpected: %s \nReceived: %s", expected, storable))
	}
}

func TestComplete(t *testing.T) {
	task := NewTask("test task")
	before := time.Now()
	task.SetCompleted(true)
	after := time.Now()
	if !(task.CompletedDate().After(before) && task.CompletedDate().Before(after)) {
		t.Error("Completed date is not set")
	}
}

func TestNotComplete(t *testing.T) {
	task := NewTask("test task")
	task.SetCompleted(true)
	task.SetCompleted(false)
	if !task.CompletedDate().IsZero() {
		t.Error("Completed date not reset")
	}
}
