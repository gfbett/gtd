package main

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

func TestPersistenceCheckFolder(t *testing.T) {

	testFolder := "testdata/wrongFolder"
	_ = os.Remove(testFolder)

	if !CreateFolderIfNotExists("testdata/existingFolder") {
		t.Fail()
	}

	if !CreateFolderIfNotExists(testFolder) {
		t.Fail()
	}

	_ = os.Remove(testFolder)
}

func TestStore(t *testing.T) {
	test := new(TestWritable)
	path := "testdata"
	if !Store(test, path) {
		t.Fail()
	}
	other := new(TestWritable)
	if !Load(other, path) {
		t.Fail()
	}
	if other.Test == false {
		t.Error("Unexpected value")
	}
}

type TestWritable struct {
	Test bool
}

func (w *TestWritable) toStorableString() string {
	return "{\"Test\":true}"
}

func (w *TestWritable) loadFromStorableString(data []byte) {
	var loaded TestWritable
	err := json.Unmarshal(data, &loaded)
	if err != nil {
		log.Fatal(err)
	}
	w.Test = loaded.Test
}

func (w *TestWritable) fileName() string {
	return "test.json"
}
