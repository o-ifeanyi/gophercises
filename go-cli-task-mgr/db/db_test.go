package db

import (
	"path/filepath"
	"strings"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
)

var home, _ = homedir.Dir()
var testPath = filepath.Join(home, "test.db")
var testBucket = []byte("test")
var completeBucket = []byte("test")

func TestInit(t *testing.T) {
	err := Init(testBucket, completeBucket, testPath)
	if err != nil {
		t.Error("DB Initialization failed:", err)
	}

	splitPath := strings.Split(db.Path(), "/")
	endsWith := splitPath[len(splitPath)-1]

	if endsWith != "test.db" {
		t.Error("Incorrect DB path:", db.Path())
	}
}

func TestCreateTask(t *testing.T) {
	_, err := CreateTask(testBucket, "Test task 1")
	if err != nil {
		t.Error("Failed to create task:", err)
	}
}

func TestAllTask(t *testing.T) {
	allTask, err := AllTasks(testBucket)
	if err != nil {
		t.Error("Failed to fetch tasks:", err)
	}
	if len(allTask) != 1 {
		t.Error("Expected to find 1 task, found", len(allTask))
	}
}

func TestDeleteTask(t *testing.T) {
	err := DeleteTask(testBucket, 1)
	if err != nil {
		t.Error("Failed to delete tasks:", err)
	}
}

func TestResetTask(t *testing.T) {
	err := ResetTask(testBucket)
	if err != nil {
		t.Error("Failed to reset tasks:", err)
	}
}

func TestItob(t *testing.T) {
	bytes := itob(1)
	if bytes[len(bytes)-1] != 1 {
		t.Error("Expected [0 0 0 0 0 0 0 1], Got", bytes)
	}
}

func TestBtoi(t *testing.T) {
	bytes := itob(1)
	i := btoi(bytes)
	if i != 1 {
		t.Error("Expected 1, Got", i)
	}
}

func BenchmarkInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Init(testBucket, completeBucket, testPath)
	}
}
