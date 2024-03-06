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
	Password    string   `json:"password,-"`
}

func main() {
	jsonData := []byte(`{
		"title": "",
		"release_year": 1972,
		"director": "juan",
		"actors": ["Marlon Brando", "Al Pacino", "James Caan", "Diane Keaton"],
		"password": "123456"
	}`)

	var movie Movie = Movie{
		Title: "Inception",
	}
	err := json.Unmarshal(jsonData, &movie)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Println(movie.Title)
	fmt.Println(movie.ReleaseYear)
	fmt.Println(movie.Director)
	fmt.Println(movie.Actors)
	fmt.Println(movie.Password)
}
