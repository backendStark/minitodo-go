package storage

import (
	"minitodo/models"
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
