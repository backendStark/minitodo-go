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
