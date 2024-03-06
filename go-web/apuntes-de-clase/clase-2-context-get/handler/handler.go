package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *MyHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "userid")

		name, ok := h.data[id]
		if !ok {
			code := http.StatusNotFound
			body := MyResponse{Message: "user not found", Data: nil}
			w.WriteHeader(code)
			w.Header().Set("Content-Type", "application/json")
			json.NewDecoder()
		}
	}
}
