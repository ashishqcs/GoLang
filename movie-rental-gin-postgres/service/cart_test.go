package service

import (
	"context"
	db "movieRentals/db/sqlc"
	"movieRentals/mocks"
	"movieRentals/model"
	"testing"
	"time"

	"go.uber.org/mock/gomock"
)

func initMockMovies() []db.Movie {
	rd := time.Now()

	return []db.Movie{
		{
			ID:       "tt0111161",
			Title:    "The Shawshank Redemption",
			Released: rd,
			Genre:    "Drama",
			Actors:   "Tim Robbins, Morgan Freeman, Bob Gunton",
			Year:     1994,
			Price:    15,
			Quantity: 4,
		},
		{
			ID:       "tt0120338",
			Title:    "Titanic",
			Released: rd,
			Genre:    "Drama, Romance",
			Actors:   "Leonardo DiCaprio, Kate Winslet, Billy Zane",
			Year:     1997,
			Price:    10,
			Quantity: 7,
		},
	}
}

func TestNewCartItemReturnsErrorWhenMovieIdIsEmpty(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStore(ctrl)

	cartItem, err := model.NewCartItem(context.Background(), "", mockStore)

	if cartItem != nil || err == nil {
		t.Errorf("expected error but got none")
	}
}

func TestNewCartReturnsSameReferenceOfCart(t *testing.T) {
	cart1 := model.NewCart()
	cart2 := model.NewCart()

	if cart1 != cart2 {
		t.Errorf("expected same but found different cart references")
	}
}

func TestAddToCartReturnsNoErrorForSingleMovieId(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStore(ctrl)
	id := "tt0111161"
	mockMovies := initMockMovies()
	movieIds := []string{id}
	mockStore.EXPECT().GetMovie(gomock.Any(), gomock.Eq(id)).Return(mockMovies[0], nil).Times(len(movieIds))

	cs := NewCartService(*model.NewCart(), mockStore)
	err := cs.AddToCart(context.Background(), movieIds)

	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}

	if cs.cart.Items[0].MovieId != id {
		t.Errorf("expected movied id %s to be in cart but found %v", id, cs.cart.Items[0].MovieId)
	}
}

func TestAddToCartReturnsNoErrorForMultipleMovieIds(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStore(ctrl)
	id1 := "tt0111161"
	id2 := "tt0120338"
	mockMovies := initMockMovies()
	movieIds := []string{id1, id2}
	mockStore.EXPECT().GetMovie(gomock.Any(), gomock.Eq(id1)).Return(mockMovies[0], nil).Times(1)
	mockStore.EXPECT().GetMovie(gomock.Any(), gomock.Eq(id2)).Return(mockMovies[1], nil).Times(1)

	cs := NewCartService(*model.NewCart(), mockStore)
	err := cs.AddToCart(context.Background(), movieIds)

	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}

	if len(cs.cart.Items) != len(movieIds) {
		t.Errorf("expected %d cart items but got %d", len(movieIds), len(cs.cart.Items))
	}

	if cs.cart.Items[0].MovieId != id1 {
		t.Errorf("expected movied id %s to be in cart but found %v", id1, cs.cart.Items[0].MovieId)
	}
}

func TestAddToCartCalculatesCorrectTotalPriceForMultipleMovieIds(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStore(ctrl)
	id1 := "tt0111161"
	id2 := "tt0120338"
	mockMovies := initMockMovies()
	movieIds := []string{id1, id2}
	expectedPrice := mockMovies[0].Price + mockMovies[1].Price
	mockStore.EXPECT().GetMovie(gomock.Any(), gomock.Eq(id1)).Return(mockMovies[0], nil).Times(1)
	mockStore.EXPECT().GetMovie(gomock.Any(), gomock.Eq(id2)).Return(mockMovies[1], nil).Times(1)

	cs := NewCartService(*model.NewCart(), mockStore)
	cs.AddToCart(context.Background(), movieIds)

	if expectedPrice != cs.cart.TotalPrice {
		t.Errorf("expected total price to be %d but got %d", expectedPrice, cs.cart.TotalPrice)
	}
}

func TestGetCartShouldReturnCartResponse(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStore(ctrl)
	id1 := "tt0111161"
	title1 := "The Shawshank Redemption"
	cart := model.Cart{
		Items: []model.CartItem{
			{
				MovieId: id1,
				Title:   title1,
			},
		},
	}
	cs := NewCartService(cart, mockStore)

	cartResponse := cs.GetCart(context.Background())

	if cartResponse.Items[0].MovieId != id1 {
		t.Errorf("expected %s but got %s", id1, cartResponse.Items[0].MovieId)
	}
	if cartResponse.Items[0].Title != title1 {
		t.Errorf("expected %s but got %s", title1, cartResponse.Items[0].Title)
	}
}

func TestGetCartShouldReturnEmptyCartInCaseOfEmptyCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStore(ctrl)

	cs := NewCartService(*model.NewCart(), mockStore)

	cartResponse := cs.GetCart(context.Background())

	if len(cartResponse.Items) != 0 {
		t.Errorf("expected size to be 0 but got %d", len(cartResponse.Items))
	}
}
