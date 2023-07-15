package service

import (
	"context"
	"errors"
	db "movieRentals/db/sqlc"
	"movieRentals/model"
)

var ErrNoMovieIdsPresent = errors.New("movie ids is empty")

type ICartService interface {
	AddToCart(ctx context.Context, movieIds []string) error
	GetCart(ctx context.Context) *model.GetCartResponse
}

type CartService struct {
	cart model.Cart
	db   db.Store
}

func NewCartService(cart model.Cart, db db.Store) *CartService {
	return &CartService{
		cart,
		db,
	}
}

func (cs *CartService) AddToCart(ctx context.Context, movieIds []string) error {
	if len(movieIds) == 0 {
		return ErrNoMovieIdsPresent
	}
	for _, movieId := range movieIds {
		cartMovie, err := model.NewCartItem(ctx, movieId, cs.db)
		if err != nil {
			return err
		}
		cs.cart.AddCartItem(cartMovie)
	}

	return nil
}

func (cs *CartService) GetCart(ctx context.Context) *model.GetCartResponse {
	return model.CartToCartResponse(&cs.cart)
}
