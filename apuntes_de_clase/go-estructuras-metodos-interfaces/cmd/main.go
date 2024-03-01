package cmd

import (
	"go-estructuras-metodos-interfaces/internal"
	"go-estructuras-metodos-interfaces/internal/service"
)

func main() {
	var rp internal.MovieRepository
	rp =
	sv := service.NewMovieDefault(rp)

	avg, err := sv.AverageRating()
}
