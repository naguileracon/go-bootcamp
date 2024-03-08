package internal

import "errors"

type Task struct {
	ID          int
	Title       string
	Description string
	Done        bool
}

var (
	ErrTaskDuplicated   = errors.New("task duplicated")
	ErrTaskInvalidField = errors.New("task invalid field")
	ErrTaskConflict     = errors.New("task conflict")
)

type TaskRepository interface {
	Save(task *Task) (err error)
}
type TaskService interface {
	Save(task *Task) (err error)
}
