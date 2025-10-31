package storage

import (
	"minitodo/models"
	"os"
	"path/filepath"
	"testing"
)

func createMockTaskManager(t *testing.T, tasks []models.Task) (*TaskManager, *MockStorage) {
	t.Helper()

	mockStorage := NewMockStorage(tasks)
	tm, err := NewTaskManagerWithStorage(mockStorage)

	if err != nil {
		t.Fatalf("Failed to create TaskManager: %v", err)
	}

	return tm, mockStorage
}

func TestTaskManager_Add_Success(t *testing.T) {
	tm, mockStorage := createMockTaskManager(t, []models.Task{})

	err := tm.Add("Buy milk")

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if tm.GetCount() != 1 {
		t.Errorf("Expected count 1, got: %d", tm.GetCount())
	}

	tasks := tm.GetAll()
	if tasks[0].Text != "Buy milk" {
		t.Errorf("Expected text 'Buy milk', got: %s", tasks[0].Text)
	}

	if tasks[0].Done != false {
		t.Errorf("Expected Done=false, got: %v", tasks[0].Done)
	}

	if mockStorage.GetSaveCalls() != 1 {
		t.Errorf("Expected 1 Save call, got: %d", mockStorage.GetSaveCalls())
	}
}

func TestTaskManager_Add_EmptyTask(t *testing.T) {
	tm, _ := createMockTaskManager(t, []models.Task{})

	err := tm.Add("")

	if err == nil {
		t.Error("Expected error for empty task, got nil")
	}

	err = tm.Add("   ")

	if err == nil {
		t.Error("Expected error for empty whitespace task, got nil")
	}

	err = tm.Add(" \n  ")

	if err == nil {
		t.Error("Expected error for whitespace with newline, got nil")
	}

	if tm.GetCount() != 0 {
		t.Errorf("Expected 0 tasks after failed add, got %d", tm.GetCount())
	}
}

func TestTaskManager_Toggle_Success(t *testing.T) {
	tm, mockStorage := createMockTaskManager(t, []models.Task{{Text: "Buy milk", Done: false}})

	err := tm.Toggle(0)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	tasks := tm.GetAll()
	if tasks[0].Done != true {
		t.Errorf("Expected Done=true, got: %v", tasks[0].Done)
	}

	if mockStorage.GetSaveCalls() != 1 {
		t.Errorf("Expected 1 Save call, got: %d", mockStorage.GetSaveCalls())
	}

	err = tm.Toggle(0)
	if err != nil {
		t.Errorf("Expected no error on second toggle, got: %v", err)
	}

	tasks = tm.GetAll()
	if tasks[0].Done != false {
		t.Errorf("Expected Done=false, got: %v", tasks[0].Done)
	}

	if mockStorage.GetSaveCalls() != 2 {
		t.Errorf("Expected 2 Save call, got: %d", mockStorage.GetSaveCalls())
	}
}

func TestTaskManager_Toggle_InvalidIndex(t *testing.T) {
	tm, _ := createMockTaskManager(t, []models.Task{{Text: "Buy milk", Done: false}})

	err := tm.Toggle(-1)

	if err == nil {
		t.Error("Expected error for negative index, got nil")
	}

	err = tm.Toggle(1)

	if err == nil {
		t.Error("Expected error for too large index, got nil")
	}
}

func TestTaskManager_Delete_Success(t *testing.T) {
	tm, mockStorage := createMockTaskManager(t, []models.Task{{Text: "Buy milk", Done: false}})

	err := tm.Delete(0)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if mockStorage.GetSaveCalls() != 1 {
		t.Errorf("Expected 1 Save call, got: %d", mockStorage.GetSaveCalls())
	}

	if tm.GetCount() != 0 {
		t.Errorf("Expected 0 tasks, got: %d", tm.GetCount())
	}
}

func TestTaskManager_Delete_InvalidIndex(t *testing.T) {
	tm, _ := createMockTaskManager(t, []models.Task{{Text: "Buy milk", Done: false}})

	err := tm.Delete(-1)

	if err == nil {
		t.Error("Expected error for negative index for delete, got nil")
	}

	err = tm.Delete(1)

	if err == nil {
		t.Error("Expected error for too large index for delete, got nil")
	}
}

func TestTaskManager_Delete_MiddleTask(t *testing.T) {
	tasks := []models.Task{
		{Text: "Task 1", Done: false},
		{Text: "Task 2", Done: false},
		{Text: "Task 3", Done: false},
	}
	tm, _ := createMockTaskManager(t, tasks)

	err := tm.Delete(1)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	remaining := tm.GetAll()

	if len(remaining) != 2 {
		t.Errorf("Expected 2 tasks, got: %d", len(remaining))
	}

	if remaining[0].Text != "Task 1" {
		t.Errorf("Expected first tasks text is 'Task 1', got: %s", remaining[0].Text)
	}

	if remaining[1].Text != "Task 3" {
		t.Errorf("Expected first tasks text is 'Task 3', got: %s", remaining[1].Text)
	}
}

func TestTaskManager_NewTaskManager_Succes(t *testing.T) {
	tmpDir := t.TempDir()
	filename := filepath.Join(tmpDir, "test_todos.json")

	tm, err := NewTaskManager(filename)

	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if tm == nil {
		t.Fatal("Expected TaskManager, got nil")
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Error("Expected file to be created")
	}

	if tm.GetCount() != 0 {
		t.Errorf("Expected 0 tasks, got: %d", tm.GetCount())
	}
}
