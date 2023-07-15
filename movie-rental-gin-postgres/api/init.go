package api

import (
	"context"
	"log"
	db "movieRentals/db/sqlc"
)

func (s *Server) load() error {
	movies, err := s.reader.GetMovies()
	if err != nil {
		return err
	}

	if len(movies.Movies) == 0 {
		return nil
	}

	for _, movie := range movies.Movies {
		args := db.CreateMovieParams{
			ID:       movie.ID,
			Title:    movie.Title,
			Released: movie.Released.Time,
			Genre:    movie.Genre,
			Actors:   movie.Actors,
			Year:     movie.Year,
			Price:    movie.Price,
			Quantity: movie.Quantity,
		}

		_, err := s.store.CreateMovie(context.Background(), args)

		if err != nil {
			log.Printf("error loading movie %v to database : %v", movie, err)
		}
	}

	return nil
}
