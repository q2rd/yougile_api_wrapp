package yougile_api_wrapp

import (
	"fmt"
	"time"
)

const taskPath = "tasks"

type Pagination struct {
	Count  int  `json:"count"`
	Limit  int  `json:"limit"`
	Offset int  `json:"offset"`
	Next   bool `json:"next"`
}
type TaskListResponse struct {
	Pagination Pagination `json:"paging"`
	Content    []*Task    `json:"content"`
}

type Deadline struct {
	Deadline  time.Duration `json:"deadline"`
	StartDate time.Duration `json:"startDate,omitempty"`
	WithTime  bool          `json:"withTime,omitempty"`
}

type TimeTracking struct {
	Plan int64 `json:"plan"`
	Work int64 `json:"work"`
}
type Checklists struct {
	Title string       `json:"title,omitempty"`
	Items []CheckItems `json:"items,omitempty"`
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
	client       *YouGileClient
	Id           string            `json:"id,omitempty"`
	Title        string            `json:"title,omitempty"`
	ColumnID     string            `json:"columnId,omitempty"`
	Description  string            `json:"description,omitempty"`
	Archived     bool              `json:"archived,omitempty"`
	Completed    bool              `json:"completed,omitempty"`
	Subtasks     []string          `json:"subtasks,omitempty"`
	Assigned     []string          `json:"assigned,omitempty"`
	Deadline     *Deadline         `json:"deadline,omitempty"`
	TimeTracking *TimeTracking     `json:"timeTracking,omitempty"`
	Checklists   *Checklists       `json:"checklists,omitempty"`
	Stickers     map[string]string `json:"stickers,omitempty"`
	Stopwatch    *Stopwatch        `json:"stopwatch,omitempty"`
	Deleted      bool              `json:"deleted,omitempty"`
}

func (c *YouGileClient) CreateTask(task *Task, extraKwargs ...Arguments) error {
	kwargs := Defaults()
	kwargs.flatten(extraKwargs)
	err := c.Post(taskPath, kwargs, &task)
	if err != nil {
		return fmt.Errorf("error creating task: %w", err)
	}
	return err
}

func (c *YouGileClient) DeleteMultiTask(tasks []*Task, extraKwargs ...Arguments) error { // TODO: подумать как лучше реализовать мультиудаление
	for _, v := range tasks {
		kwargs := Defaults()
		kwargs.flatten(extraKwargs)
		pathWithId := fmt.Sprintf("%s/%s", taskPath, v.Id)
		t := &Task{Deleted: true}
		err := c.Put(pathWithId, kwargs, &t)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Column) GetTaskList() (tasks *TaskListResponse, err error) {
	var kwargs Arguments
	if c.Id != "" {
		kwargs = Arguments{
			"columnId": c.Id,
		}
	} else {
		return nil, fmt.Errorf("columnId is required")
	}
	var response TaskListResponse
	err = c.Get(taskPath, kwargs, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (c *YouGileClient) DeleteTask(task *Task) error {
	kwargs := Defaults()
	pathWithId := fmt.Sprintf("%s/%s", taskPath, task.Id)
	t := &Task{Deleted: true}
	err := c.Put(pathWithId, kwargs, &t)
	if err != nil {
		return err // TODO логировние
	}
	return err
}
