package main

import "fmt"

type Movie struct {
	movieId  int
	title    string
	duration int
	genre    string
	language string
	rating   float64
}

func (m *Movie) GetDetails() {
	fmt.Printf("Movie: %s\nDuration: %d minutes\nGenre: %s\nLanguage: %s\nRating: %.1f\n", m.title, m.duration, m.genre, m.language, m.rating)

}

func (m *Movie) GetShows() {
	fmt.Println("Getting showtimes for movie...")
}
