package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	/// handler
	handlerPing := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("pong"))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	// router
	router := chi.NewRouter()
	router.Get("/ping", handlerPing)

	// start the server
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
