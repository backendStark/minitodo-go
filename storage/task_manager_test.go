package storage

import (
	"minitodo/models"
	"testing"
)

func TestTaskManager_Add_Success(t *testing.T) {
	mockStorage := NewMockStorage([]models.Task{})
	tm, err := NewTaskManagerWithStorage(mockStorage)

	if err != nil {
		t.Fatalf("Failed to create TaskManager: %v", err)
	}

	err = tm.Add("Buy milk")

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
