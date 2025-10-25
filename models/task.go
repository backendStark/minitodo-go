package models

type Task struct {
	Text string `json:"text"`
	Done bool   `json:"done"`
}
