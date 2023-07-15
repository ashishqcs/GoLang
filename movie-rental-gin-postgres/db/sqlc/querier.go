// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0

package db

import (
	"context"
)

type Querier interface {
	CreateMovie(ctx context.Context, arg CreateMovieParams) (Movie, error)
	GetMovie(ctx context.Context, id string) (Movie, error)
	ListMovies(ctx context.Context, arg ListMoviesParams) ([]Movie, error)
}

var _ Querier = (*Queries)(nil)
