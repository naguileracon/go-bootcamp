package service

import (
	"app/internal"
)

func NewTaskDefault(rp internal.TaskRepository) *TaskDefault {
	return &TaskDefault{rp: rp}
}

type TaskDefault struct {
	rp internal.TaskRepository
}

func (t *TaskDefault) Save(task *internal.Task) (err error) {
	err = t.rp.Save(task)
	return
}
