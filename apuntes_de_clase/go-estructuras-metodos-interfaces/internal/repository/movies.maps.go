package repository

import (
	"go-estructuras-metodos-interfaces/internal"
)

var movies []internal.Movie

type MoviesSlice struct {
	movies []internal.Movie
	// limit int
	// count int
}

func (m *MoviesSlice) Getno() (mv []internal.Movie) {
	mv = make([]internal.Movie, len(movies))
	copy(mv, movies)
	return
}
