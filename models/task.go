package models

type Task struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Assignee string `json:"assignee"`
	Priority string `json:"priority"`
	Status   string `json:"status"`
	DueDate  string `json:"dueDate"`
}
