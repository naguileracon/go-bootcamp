package main

import (
	"app/internal/handlers"
	"app/internal/repository"
	"app/internal/service"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	// dependencies
	// - repository
	rp := repository.NewProductMap(nil, 0)
	// - service
	sv := service.NewServiceDefault(rp)
	// - handler
	hd := handlers.NewDefaultProduct(sv)
	// - router
	router := chi.NewRouter()
	router.Route("/products", func(r chi.Router) {
		// POST /tasks
		r.Post("/", hd.Create())
		r.Get("/{id}", hd.GetById())
	})

	// server
	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		fmt.Println(err)
		return
	}
}
