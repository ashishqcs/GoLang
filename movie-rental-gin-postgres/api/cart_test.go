package api

import (
	"bytes"
	"encoding/json"
	"movieRentals/mocks"
	"movieRentals/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
)

func TestAddToCartReturnsNoErrorWhenOneOrMoreMovieIdsAreGiven(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStore(ctrl)
	mockReader := mocks.NewMockMoviesReader(ctrl)
	mockCartService := mocks.NewMockICartService(ctrl)

	mockCartService.EXPECT().AddToCart(gomock.Any(), gomock.Any()).Return(nil).Times(1)

	mockPriceService := mocks.NewMockIPriceService(ctrl)

	gin.SetMode(gin.TestMode)
	server, err := NewServer(mockStore, mockReader, mockCartService, mockPriceService)
	if err != nil {
		t.Errorf("expected no error while creating server, but got %v", err)
	}

	movieIds := []string{"test1", "test2"}

	req := addToCartRequest{
		MovieIds: movieIds,
	}
	body, _ := json.Marshal(req)
	recorder := httptest.NewRecorder()
	url := "/cart/movies"
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		t.Errorf("expected no error while crearting request, but got %v", err)
	}

	server.router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusCreated {
		t.Errorf("expected status code 201 but got %d", recorder.Code)
	}
}

func TestGetCartReturnsOkStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStore(ctrl)
	mockReader := mocks.NewMockMoviesReader(ctrl)
	mockCartService := mocks.NewMockICartService(ctrl)

	mockCartService.EXPECT().GetCart(gomock.Any()).Return(&model.GetCartResponse{}).Times(1)

	mockPriceService := mocks.NewMockIPriceService(ctrl)

	gin.SetMode(gin.TestMode)
	server, err := NewServer(mockStore, mockReader, mockCartService, mockPriceService)
	if err != nil {
		t.Errorf("expected no error while creating server, but got %v", err)
	}

	recorder := httptest.NewRecorder()
	url := "/cart/movies"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("expected no error while crearting request, but got %v", err)
	}

	server.router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code 200 but got %d", recorder.Code)
	}
}
