package models

type Task struct {
	Text string `json:"text"`
	Done bool   `json:"done"`
}

func (t *Task) Toggle() {
	t.Done = !t.Done
}

type SortMode int

const (
	SortByStatus SortMode = iota
	SortByStatusReverse
)

func (sm SortMode) String() string {
	switch sm {
	case SortByStatus:
		return "↑Active first"
	case SortByStatusReverse:
		return "↓Done first"
	default:
		return "Unknown"
	}
}
