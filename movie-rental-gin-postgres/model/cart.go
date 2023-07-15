package model

import (
	"context"
	"database/sql"
	"errors"
	db "movieRentals/db/sqlc"
	"sync"
)

var ErrMovieIdNotFound = errors.New("movie id not found")
var ErrInvalidMovieId = errors.New("invalid movie id")

var cartInstance Cart

type CartItem struct {
	MovieId string
	Price   int64
	Title   string
}

func NewCartItem(ctx context.Context, movieId string, db db.Store) (*CartItem, error) {
	if movieId == "" {
		return nil, ErrInvalidMovieId
	}

	movie, err := db.GetMovie(ctx, movieId)

	if errors.Is(err, sql.ErrNoRows) {

		return nil, ErrMovieIdNotFound
	}

	if err != nil {
		return nil, err
	}

	return &CartItem{
		MovieId: movieId,
		Price:   movie.Price,
		Title:   movie.Title,
	}, nil

}

type Cart struct {
	Items      []CartItem
	TotalPrice int64
}

func NewCart() *Cart {
	var s sync.Once

	s.Do(
		func() {
			cartInstance = Cart{
				Items:      []CartItem{},
				TotalPrice: 0,
			}
		},
	)
	return &cartInstance
}

func (c *Cart) AddCartItem(movieToBeAdded *CartItem) {
	for _, movie := range c.Items {
		if movie == *movieToBeAdded {
			return
		}
	}

	c.Items = append(c.Items, *movieToBeAdded)
	c.TotalPrice += movieToBeAdded.Price
}

type GetCartResponse struct {
	Items []CartItemResponse
}

type CartItemResponse struct {
	MovieId string
	Title   string
}

func CartToCartResponse(cart *Cart) *GetCartResponse {
	items := []CartItemResponse{}
	for _, cartItem := range cart.Items {
		responseItem := CartItemResponse{
			MovieId: cartItem.MovieId,
			Title:   cartItem.Title,
		}

		items = append(items, responseItem)
	}

	return &GetCartResponse{
		Items: items,
	}
}
