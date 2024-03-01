package service

import "go-estructuras-metodos-interfaces/internal"

type MovieDefault struct {
	rp internal.MovieRepository
}

func (m MovieDefault) AverageRating() (avg float64, err string) {
	movies := m.rp.Get()

	if len(movies) == 0 {
		err = "no movies found"
		return
	}
	var sum float64
	for _, movie := range movies {
		sum += movie.Rating
	}

	avg = sum / float64(len(movies))
	return
}
