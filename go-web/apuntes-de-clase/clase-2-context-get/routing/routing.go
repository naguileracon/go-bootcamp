package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("users"))
		if err != nil {
			println(err.Error())
			return
		}
	})

	// print the id of the user
	r.Get("/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("user by id"))
		if err != nil {
			println(err.Error())
			return
		}
	})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		println(err.Error())
		return
	}

}
