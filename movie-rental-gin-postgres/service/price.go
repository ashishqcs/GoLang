package service

import (
	"context"
	"errors"
	"fmt"
	db "movieRentals/db/sqlc"
	"movieRentals/model"
	"movieRentals/strategy"
	"time"
)

var ErrInvalidDates = errors.New("rentFrom date should be before rentTo date")
var dateFormat = "2006-01-02"

type IPriceService interface {
	PriceItem(ctx context.Context, items []model.PricingItem) (*model.PricingResponse, error)
}

type PriceService struct {
	db db.Store
}

func NewPriceService(db db.Store) *PriceService {
	return &PriceService{
		db,
	}
}

func (ps *PriceService) PriceItem(ctx context.Context, items []model.PricingItem) (*model.PricingResponse, error) {
	responseItems := []model.PricingItemResponse{}
	var totalPrice int64

	for _, item := range items {
		from, _ := time.Parse(dateFormat, item.RentFrom)
		to, _ := time.Parse(dateFormat, item.RentTo)

		if from.After(to) {
			return nil, ErrInvalidDates
		}

		movie, err := ps.db.GetMovie(ctx, item.MovieId)
		if err != nil {
			return nil, err
		}

		pricedItem := price(movie, from, to)

		responseItems = append(responseItems, *pricedItem)
		totalPrice += pricedItem.MoviePrice
	}
	return model.NewPricingResponse(responseItems, totalPrice), nil
}

func price(movie db.Movie, from time.Time, to time.Time) *model.PricingItemResponse {
	days := int64(to.Sub(from).Hours() / 24)

	price := applyPricingStrategy(movie, days)

	duration := fmt.Sprintf("%s to %s", from.Format(dateFormat), to.Format(dateFormat))
	return model.NewPricingItemResponse(movie.ID, price, duration)
}

func applyPricingStrategy(movie db.Movie, days int64) int64 {
	var pricingStrategy strategy.PricingStrategy

	if movie.Year >= 1995 {
		pricingStrategy = strategy.NewNewMoviePricingStrategy()

	} else {
		pricingStrategy = strategy.NewClassicMoviePricingStrategy()
	}

	return pricingStrategy.PriceItem(days)
}
