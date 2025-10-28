package models

type Task struct {
	Text string `json:"text"`
	Done bool   `json:"done"`
}

func (t *Task) Toggle() {
	t.Done = !t.Done
}
