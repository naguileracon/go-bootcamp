package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title       string   `json:"title"`
	ReleaseYear int      `json:"release_year"`
	Director    string   `json:"director,omitempty"`
	Actors      []string `json:"actors"`
	Password    string   `json:"-"`
}

func main() {
	movie := Movie{
		Title:       "The Godfather",
		ReleaseYear: 1972,
		Director:    "",
		Actors:      []string{"Marlon Brando", "Al Pacino", "James Caan", "Diane Keaton"},
	}
	bytes, err := json.Marshal(movie)
	if err != nil {
		println(err.Error())
		return
	}

	fmt.Println(string(bytes))
}
