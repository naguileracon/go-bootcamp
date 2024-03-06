package internal

import "time"

type MovieAttributes struct {
	title    string
	Year     int
	Director string
	Genre    string
	Duration time.Duration
	Rating   float64
}

type Movie struct {
	ID int
	MovieAttributes
}
