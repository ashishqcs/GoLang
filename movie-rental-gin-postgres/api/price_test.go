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

func TestFindPriceReturnsNoErrorForValidRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStore(ctrl)
	mockReader := mocks.NewMockMoviesReader(ctrl)
	mockCartService := mocks.NewMockICartService(ctrl)
	mockPriceService := mocks.NewMockIPriceService(ctrl)

	gin.SetMode(gin.TestMode)
	server, err := NewServer(mockStore, mockReader, mockCartService, mockPriceService)
	if err != nil {
		t.Errorf("expected no error while creating server, but got %v", err)
	}

	req := pricingRequest{
		Items: []pricingItemRequest{
			{"abc", "2023-07-13", "2023-07-15"}, {"pqr", "2023-07-13", "2023-07-15"},
		},
	}

	response := model.PricingResponse{}
	mockPriceService.EXPECT().PriceItem(gomock.Any(), gomock.Any()).Return(&response, nil).Times(1)

	body, _ := json.Marshal(req)
	recorder := httptest.NewRecorder()
	url := "/price"
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		t.Errorf("expected no error while crearting request, but got %v", err)
	}

	server.router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code 200 but got %d", recorder.Code)
	}
}

func TestFindPriceReturnsStatusBadRequestWhenInvalidDatesAreProvided(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStore(ctrl)
	mockReader := mocks.NewMockMoviesReader(ctrl)
	mockCartService := mocks.NewMockICartService(ctrl)
	mockPriceService := mocks.NewMockIPriceService(ctrl)

	gin.SetMode(gin.TestMode)
	server, err := NewServer(mockStore, mockReader, mockCartService, mockPriceService)

	if err != nil {
		t.Errorf("expected no error while creating server, but got %v", err)
	}

	req := pricingRequest{
		Items: []pricingItemRequest{
			{"abc", "2023-07-131", "2023-07-15"}, {"pqr", "2023-07-13", "2023-07-15"},
		},
	}

	body, _ := json.Marshal(req)
	recorder := httptest.NewRecorder()
	url := "/price"
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		t.Errorf("expected no error while crearting request, but got %v", err)
	}

	server.router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusBadRequest {
		t.Errorf("expected status code 400 but got %d", recorder.Code)
	}
}
