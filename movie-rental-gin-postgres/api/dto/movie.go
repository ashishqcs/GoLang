package dto

import (
	db "movieRentals/db/sqlc"
	"time"
)

type Movie struct {
	ID       string
	Title    string
	Released time.Time
	Genre    string
	Actors   string
	Price    int64
}

func MapMovieFromDbModel(m *db.Movie) *Movie {
	return &Movie{
		ID:       m.ID,
		Title:    m.Title,
		Genre:    m.Genre,
		Released: m.Released,
		Actors:   m.Actors,
		Price:    m.Price,
	}
}

func MapMoviesFromDbModel(m []db.Movie) []Movie {
	var movies []Movie
	for _, movie := range m {
		movies = append(movies, *MapMovieFromDbModel(&movie))
	}
	return movies
}
