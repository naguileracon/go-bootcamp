package clase_2_context_get

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	h := handler.NewHandler()

	r.Get("/", h.Get())
	r.Get("/users/{userid}", h.GetByID())

	r.Get("/users", h.GetByQuery())

	http.ListenAndServe(":8080", r)
}
