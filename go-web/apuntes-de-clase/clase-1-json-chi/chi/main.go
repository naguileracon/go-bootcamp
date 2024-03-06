package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	handlerHome := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello, World!"))
		if err != nil {
			println(err.Error())
			return
		}
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
	})
	handlerHealth := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("OK"))
		if err != nil {
			println(err.Error())
			return
		}
		writer.Header().Add("Content-Type", "text/plain")
		writer.WriteHeader(http.StatusOK)
	})
	router := chi.NewRouter()
	router.Get("/", handlerHome)
	router.Get("/health", handlerHealth)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		println(err.Error())
		return
	}

	/*err := http.ListenAndServe(":8080", nil)
	if err != nil {
		println(err.Error())
		return
	}*/
	/*handler := func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello, World!"))
		if err != nil {
			println(err.Error())
			return
		}
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
	}

	err := http.ListenAndServe(":8080", http.HandlerFunc(handler))
	if err != nil {
		println(err.Error())
		return
	}*/
}
