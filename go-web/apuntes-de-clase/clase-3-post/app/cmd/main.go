package main

import (
	task "app/internal"
	"app/internal/handlers"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {

	handler := handlers.NewDefaultTask(
		map[int]task.Task{},
		0,
	)
	router := chi.NewRouter()

	router.Post("/tasks", handler.Create())

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
