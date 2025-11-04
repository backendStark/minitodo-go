package models

import "testing"

func TestTaskModel_Toggle_Success(t *testing.T) {
	task := Task{Text: "Buy milk", Done: false}
	task.Toggle()

	if task.Done != true {
		t.Error("Expect Done equal true, but got:", task.Done)
	}

	task.Toggle()

	if task.Done != false {
		t.Error("Expect Done equal false, but got:", task.Done)
	}
}

func TestSortMode_String(t *testing.T) {
	tests := []struct {
		mode     SortMode
		expected string
	}{
		{SortByStatus, "↑Active first"},
		{SortByStatusReverse, "↓Done first"},
		{SortMode(999), "Unknown"},
	}

	for _, tt := range tests {
		got := tt.mode.String()
		if got != tt.expected {
			t.Errorf("SortMode(%d).String() = %q, expected %q", tt.mode, got, tt.expected)
		}
	}
}
