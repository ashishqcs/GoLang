package service

import (
	"context"
	mocks "movieRentals/mocks"
	"movieRentals/model"
	"testing"
	"time"

	"go.uber.org/mock/gomock"
)

func TestShouldReturnOneForFiveRentalDaysOfClassicMovies(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStore(ctrl)
	id := "tt0111161"
	mockMovies := initMockMovies()

	mockStore.EXPECT().GetMovie(gomock.Any(), gomock.Eq(id)).Return(mockMovies[0], nil).Times(1)

	rentFrom := time.Now()
	rentTill := rentFrom.AddDate(0, 0, 5).Format("2006-01-02")

	ps := NewPriceService(mockStore)

	items := []model.PricingItem{
		{
			MovieId:  id,
			RentFrom: rentFrom.Format("2006-01-02"),
			RentTo:   rentTill,
		},
	}

	pricingResponse, err := ps.PriceItem(context.Background(), items)

	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}

	if pricingResponse.TotalPrice != 1 {
		t.Errorf("expected %d but found %d", 1, pricingResponse.TotalPrice)
	}
	if pricingResponse.Movies[0].MoviePrice != 1 {
		t.Errorf("expected 5 but found %d", pricingResponse.Movies[0].MoviePrice)
	}
	expectedString := items[0].RentFrom + " to " + items[0].RentTo
	if pricingResponse.Movies[0].RentDuration != expectedString {
		t.Errorf("expected %s but found %s", expectedString, pricingResponse.Movies[0].RentDuration)
	}
}

func TestShouldReturnSixForTwelveRentalDaysOfClassicMovies(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStore(ctrl)
	id := "tt0111161"
	mockMovies := initMockMovies()

	mockStore.EXPECT().GetMovie(gomock.Any(), gomock.Eq(id)).Return(mockMovies[0], nil).Times(1)

	rentFrom := time.Now()
	rentTill := rentFrom.AddDate(0, 0, 12).Format("2006-01-02")

	ps := NewPriceService(mockStore)

	items := []model.PricingItem{
		{
			MovieId:  id,
			RentFrom: rentFrom.Format("2006-01-02"),
			RentTo:   rentTill,
		},
	}

	pricingResponse, err := ps.PriceItem(context.Background(), items)

	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}

	if pricingResponse.TotalPrice != 6 {
		t.Errorf("expected %d but found %d", 6, pricingResponse.TotalPrice)
	}
	if pricingResponse.Movies[0].MoviePrice != 6 {
		t.Errorf("expected 5 but found %d", pricingResponse.Movies[0].MoviePrice)
	}
	expectedString := items[0].RentFrom + " to " + items[0].RentTo
	if pricingResponse.Movies[0].RentDuration != expectedString {
		t.Errorf("expected %s but found %s", expectedString, pricingResponse.Movies[0].RentDuration)
	}
}

func TestShouldReturnCorrectPricingForNewMovies(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStore(ctrl)
	id := "tt0120338"
	mockMovies := initMockMovies()

	mockStore.EXPECT().GetMovie(gomock.Any(), gomock.Eq(id)).Return(mockMovies[1], nil).Times(1)

	rentFrom := time.Now()
	rentTill := rentFrom.AddDate(0, 0, 5).Format("2006-01-02")

	ps := NewPriceService(mockStore)

	items := []model.PricingItem{
		{
			MovieId:  id,
			RentFrom: rentFrom.Format("2006-01-02"),
			RentTo:   rentTill,
		},
	}

	pricingResponse, err := ps.PriceItem(context.Background(), items)

	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}

	if pricingResponse.TotalPrice != 9 {
		t.Errorf("expected 9 but found %d", pricingResponse.TotalPrice)
	}
	if pricingResponse.Movies[0].MoviePrice != 9 {
		t.Errorf("expected 9 but found %d", pricingResponse.Movies[0].MoviePrice)
	}
	expectedString := items[0].RentFrom + " to " + items[0].RentTo
	if pricingResponse.Movies[0].RentDuration != expectedString {
		t.Errorf("expected %s but found %s", expectedString, pricingResponse.Movies[0].RentDuration)
	}
}

func TestShouldReturnPricingForCombOfNewAndClassicMovies(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStore(ctrl)
	id1 := "tt0111161"
	id2 := "tt0120338"
	mockMovies := initMockMovies()

	cart := model.NewCart()
	cart.Items = []model.CartItem{{MovieId: id1}, {MovieId: id2}}

	mockStore.EXPECT().GetMovie(gomock.Any(), gomock.Eq(id1)).Return(mockMovies[0], nil).Times(1)
	mockStore.EXPECT().GetMovie(gomock.Any(), gomock.Eq(id2)).Return(mockMovies[1], nil).Times(1)

	rentFrom := time.Now().AddDate(0, 0, 2)
	rentTill := rentFrom.AddDate(0, 0, 5).Format("2006-01-02")

	ps := NewPriceService(mockStore)

	items := []model.PricingItem{
		{
			MovieId:  id1,
			RentFrom: rentFrom.Format("2006-01-02"),
			RentTo:   rentTill,
		},
		{
			MovieId:  id2,
			RentFrom: rentFrom.Format("2006-01-02"),
			RentTo:   rentTill,
		},
	}

	pricingResponse, err := ps.PriceItem(context.Background(), items)

	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}

	if pricingResponse.TotalPrice != 10 {
		t.Errorf("expected 10 but found %d", pricingResponse.TotalPrice)
	}
}

func TestShouldReturnErrorForInvalidDates(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStore(ctrl)
	id := "tt0111161"

	mockStore.EXPECT().GetMovie(gomock.Any(), gomock.Eq(id)).Times(0)

	rentFrom := time.Now().AddDate(0, 0, 2)
	rentTill := rentFrom.AddDate(0, 0, -1).Format("2006-01-02")

	ps := NewPriceService(mockStore)

	items := []model.PricingItem{
		{
			MovieId:  id,
			RentFrom: rentFrom.Format("2006-01-02"),
			RentTo:   rentTill,
		},
	}

	pricingResponse, err := ps.PriceItem(context.Background(), items)

	if err == nil || pricingResponse != nil {
		t.Errorf("expected an error but got no error")
	}

	if err != ErrInvalidDates {
		t.Errorf("expected %v but found %v", ErrInvalidDates, err)
	}
}
