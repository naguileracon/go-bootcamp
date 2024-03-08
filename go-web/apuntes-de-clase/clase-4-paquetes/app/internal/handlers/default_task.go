package handlers

import (
	"app/internal"
)

func newDefaultTask() *DefaultTask {
	return &DefaultTask{}
}

type DefaultTask struct {
	sv internal.TaskService

type TaskJSON struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type TaskRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func (d *DefaultTask) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// validate token

		token := r.Header.Get("Authorization")
		if token != "12345" {
			web.ResponseJSON(w, http.StatusUnauthorized, map[string]any{
				"message": "unauthorized",
			})
		}
		// request
		//validate require fields
		CheckRequiredFields(w, r, []string{"title", "description", "done"})
		bodyBytes, _ := io.ReadAll(r.Body)
		bodyMap := map[string]any{}
		err := json.Unmarshal(bodyBytes, &bodyMap)

		if err != nil {
			println(err.Error())
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid request body",
			})
			return
		}

		// validate required keys
		if _, ok := bodyMap["title"]; !ok {
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{
				"message": "title is required",
			})
			return
		}
		if _, ok := bodyMap["description"]; !ok {
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{
				"message": "description is required",
			})
			return
		}
		if _, ok := bodyMap["done"]; !ok {
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{
				"message": "done is required",
			})
			return
		}

		// parsing request task request body
		var body TaskRequestBody
		err = json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			println(err.Error())
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid request body",
			})
			return
		}

		//- validate the task

		// process
		d.lastID++
		createdTask := task.Task{
			ID:          d.lastID,
			Title:       body.Title,
			Description: body.Description,
			Done:        body.Done,
		}

		// validate task

		if createdTask.Title == "" || len(createdTask.Title) > 25 {
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{
				"message": "title is required",
			})
			return
		}

		d.task[createdTask.ID] = createdTask

		// response
		data := TaskJSON{
			Id:          createdTask.ID,
			Title:       createdTask.Title,
			Description: createdTask.Description,
			Done:        createdTask.Done,
		}

		web.ResponseJSON(w, http.StatusCreated, map[string]any{"message": "task created", "task": data})

	}
}
