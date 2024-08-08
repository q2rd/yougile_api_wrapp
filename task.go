package yougile_api_wrapp

import "time"

type Deadline struct {
	Deadline  time.Duration `json:"deadline"`
	StartDate time.Duration `json:"startDate"`
	WithTime  bool          `json:"withTime"`
}

type TimeTracking struct {
	Plan int64 `json:"plan"`
	Work int64 `json:"work"`
}
type Checklists struct {
	Title string       `json:"title"`
	Items []CheckItems `json:"items"`
}

type CheckItems struct {
	ChecklistTitle string `json:"title"`
	IsCompleted    bool   `json:"isCompleted"`
}

type Stopwatch struct {
	Running bool `json:"running"`
}

type Timer struct {
	Seconds time.Duration `json:"seconds"`
	Running bool          `json:"running"`
}

type Task struct {
	Column       *Column
	Title        string            `json:"title"`
	ColumnID     string            `json:"column_id"`
	Description  string            `json:"description"`
	Archived     bool              `json:"archived"`
	Completed    bool              `json:"completed"`
	Subtasks     []string          `json:"subtasks"`
	Assigned     []string          `json:"assigned"`
	Deadline     Deadline          `json:"deadline"`
	TimeTracking TimeTracking      `json:"timeTracking"`
	Checklists   Checklists        `json:"checklists"`
	Stickers     map[string]string `json:"stickers"`
	Stopwatch    Stopwatch         `json:"stopwatch"`
}
