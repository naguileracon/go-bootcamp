package repository

import "app/internal"

func NewTaskMap(db map[int]internal.Task, lastID int) *TaskMap {
	if db == nil {
		db = make(map[int]internal.Task)
	}
	return &TaskMap{db: db, lastId: lastID}
}

type TaskMap struct {
	db     map[int]internal.Task
	lastId int
}

func (t *TaskMap) Save(task *internal.Task) (err error) {
	t.lastId++
	task.ID = t.lastId
	t.db[task.ID] = *task
	return
}
