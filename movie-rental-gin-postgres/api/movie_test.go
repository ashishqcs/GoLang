package api

import (
	"database/sql"
	"errors"
	"fmt"
	db "movieRentals/db/sqlc"
	"movieRentals/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
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

func TestListMoviesShouldListAllMovies(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStore(ctrl)
	mockMovies := initMockMovies()
	mockReader := mocks.NewMockMoviesReader(ctrl)
	mockCartService := mocks.NewMockICartService(ctrl)
	mockStore.EXPECT().ListMovies(gomock.Any(), gomock.Any()).Return(mockMovies, nil).Times(1)
	mockPriceService := mocks.NewMockIPriceService(ctrl)

	gin.SetMode(gin.TestMode)
	server, err := NewServer(mockStore, mockReader, mockCartService, mockPriceService)
	if err != nil {
		t.Errorf("expected no error while creating server, but got %v", err)
	}

	recorder := httptest.NewRecorder()
	url := "/movies"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("expected no error while crearting request, but got %v", err)
	}

	q := request.URL.Query()
	q.Add("page_id", "1")
	q.Add("page_size", "2")
	request.URL.RawQuery = q.Encode()

	server.router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code 200 but got %d", recorder.Code)
	}
}

func TestListMoviesShouldReturnInternalServerErrorWhenDbCallReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStore(ctrl)
	mockReader := mocks.NewMockMoviesReader(ctrl)
	mockCartService := mocks.NewMockICartService(ctrl)
	mockStore.EXPECT().ListMovies(gomock.Any(), gomock.Any()).Return(nil, errors.New("mock db error")).Times(1)
	mockPriceService := mocks.NewMockIPriceService(ctrl)

	gin.SetMode(gin.TestMode)
	server, err := NewServer(mockStore, mockReader, mockCartService, mockPriceService)
	if err != nil {
		t.Errorf("expected no error while creating server, but got %v", err)
	}

	recorder := httptest.NewRecorder()
	url := "/movies"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("expected no error while crearting request, but got %v", err)
	}

	q := request.URL.Query()
	q.Add("page_id", "1")
	q.Add("page_size", "2")
	request.URL.RawQuery = q.Encode()

	server.router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusInternalServerError {
		t.Errorf("expected status code 500 but got %d", recorder.Code)
	}
}

func TestListMoviesShouldReturnBadRequestErrorWhenPageIdIsIncorrect(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStore(ctrl)
	mockReader := mocks.NewMockMoviesReader(ctrl)
	mockCartService := mocks.NewMockICartService(ctrl)
	mockStore.EXPECT().ListMovies(gomock.Any(), gomock.Any()).Times(0)
	mockPriceService := mocks.NewMockIPriceService(ctrl)

	gin.SetMode(gin.TestMode)
	server, err := NewServer(mockStore, mockReader, mockCartService, mockPriceService)
	if err != nil {
		t.Errorf("expected no error while creating server, but got %v", err)
	}

	recorder := httptest.NewRecorder()
	url := "/movies"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("expected no error while crearting request, but got %v", err)
	}

	q := request.URL.Query()
	q.Add("page_id", "-1")
	q.Add("page_size", "1")
	request.URL.RawQuery = q.Encode()

	server.router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusBadRequest {
		t.Errorf("expected status code 400 but got %d", recorder.Code)
	}
}

func TestListMoviesShouldReturnBadRequestErrorWhenPageSizeIsIncorrect(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStore(ctrl)
	mockReader := mocks.NewMockMoviesReader(ctrl)
	mockCartService := mocks.NewMockICartService(ctrl)
	mockStore.EXPECT().ListMovies(gomock.Any(), gomock.Any()).Times(0)
	mockPriceService := mocks.NewMockIPriceService(ctrl)

	gin.SetMode(gin.TestMode)
	server, err := NewServer(mockStore, mockReader, mockCartService, mockPriceService)
	if err != nil {
		t.Errorf("expected no error while creating server, but got %v", err)
	}

	recorder := httptest.NewRecorder()
	url := "/movies"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("expected no error while crearting request, but got %v", err)
	}

	q := request.URL.Query()
	q.Add("page_id", "1")
	q.Add("page_size", "6")
	request.URL.RawQuery = q.Encode()

	server.router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusBadRequest {
		t.Errorf("expected status code 400 but got %d", recorder.Code)
	}
}

func TestListMoviesShouldReturnMoviesFilteredByYear(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStore(ctrl)
	mockMovies := initMockMovies()
	mockReader := mocks.NewMockMoviesReader(ctrl)
	mockCartService := mocks.NewMockICartService(ctrl)
	mockStore.EXPECT().ListMovies(gomock.Any(), gomock.Any()).Return(mockMovies, nil).Times(1)
	mockPriceService := mocks.NewMockIPriceService(ctrl)

	gin.SetMode(gin.TestMode)
	server, err := NewServer(mockStore, mockReader, mockCartService, mockPriceService)
	if err != nil {
		t.Errorf("expected no error while creating server, but got %v", err)
	}

	recorder := httptest.NewRecorder()
	url := "/movies"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("expected no error while crearting request, but got %v", err)
	}

	q := request.URL.Query()
	q.Add("page_id", "1")
	q.Add("page_size", "2")
	q.Add("year", "1972")
	request.URL.RawQuery = q.Encode()

	server.router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code 200 but got %d", recorder.Code)
	}
}

func TestListMoviesShouldReturnBadRequestWhenYearIsLessThan1950(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStore(ctrl)
	mockReader := mocks.NewMockMoviesReader(ctrl)
	mockCartService := mocks.NewMockICartService(ctrl)
	mockStore.EXPECT().ListMovies(gomock.Any(), gomock.Any()).Times(0)
	mockPriceService := mocks.NewMockIPriceService(ctrl)

	gin.SetMode(gin.TestMode)
	server, err := NewServer(mockStore, mockReader, mockCartService, mockPriceService)
	if err != nil {
		t.Errorf("expected no error while creating server, but got %v", err)
	}

	recorder := httptest.NewRecorder()
	url := "/movies"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("expected no error while crearting request, but got %v", err)
	}

	q := request.URL.Query()
	q.Add("page_id", "1")
	q.Add("page_size", "2")
	q.Add("year", "1949")
	request.URL.RawQuery = q.Encode()

	server.router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusBadRequest {
		t.Errorf("expected status code 400 but got %d", recorder.Code)
	}
}

func TestListMoviesShouldReturnBadRequestWhenYearIsGreaterThanCurrentYear(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStore(ctrl)
	mockReader := mocks.NewMockMoviesReader(ctrl)
	mockCartService := mocks.NewMockICartService(ctrl)
	mockStore.EXPECT().ListMovies(gomock.Any(), gomock.Any()).Times(0)
	mockPriceService := mocks.NewMockIPriceService(ctrl)

	gin.SetMode(gin.TestMode)
	server, err := NewServer(mockStore, mockReader, mockCartService, mockPriceService)
	if err != nil {
		t.Errorf("expected no error while creating server, but got %v", err)
	}

	recorder := httptest.NewRecorder()
	url := "/movies"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("expected no error while crearting request, but got %v", err)
	}

	q := request.URL.Query()
	q.Add("page_id", "1")
	q.Add("page_size", "2")
	q.Add("year", fmt.Sprint(time.Now().Year()+1))
	request.URL.RawQuery = q.Encode()

	server.router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusBadRequest {
		t.Errorf("expected status code 400 but got %d", recorder.Code)
	}
}

func TestListMoviesShouldReturnMoviesFilteredByGenre(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStore(ctrl)
	mockMovies := initMockMovies()
	mockReader := mocks.NewMockMoviesReader(ctrl)
	mockCartService := mocks.NewMockICartService(ctrl)
	mockStore.EXPECT().ListMovies(gomock.Any(), gomock.Any()).Return(mockMovies, nil).Times(1)

	mockPriceService := mocks.NewMockIPriceService(ctrl)

	gin.SetMode(gin.TestMode)
	server, err := NewServer(mockStore, mockReader, mockCartService, mockPriceService)
	if err != nil {
		t.Errorf("expected no error while creating server, but got %v", err)
	}

	recorder := httptest.NewRecorder()
	url := "/movies"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("expected no error while crearting request, but got %v", err)
	}

	q := request.URL.Query()
	q.Add("page_id", "1")
	q.Add("page_size", "2")
	q.Add("genre", "drama")
	request.URL.RawQuery = q.Encode()

	server.router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code 200 but got %d", recorder.Code)
	}
}

func TestListMoviesShouldReturnMoviesFilteredByActor(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStore(ctrl)
	mockMovies := initMockMovies()
	mockReader := mocks.NewMockMoviesReader(ctrl)
	mockCartService := mocks.NewMockICartService(ctrl)
	mockPriceService := mocks.NewMockIPriceService(ctrl)
	mockStore.EXPECT().ListMovies(gomock.Any(), gomock.Any()).Return(mockMovies, nil).Times(1)

	gin.SetMode(gin.TestMode)
	server, err := NewServer(mockStore, mockReader, mockCartService, mockPriceService)
	if err != nil {
		t.Errorf("expected no error while creating server, but got %v", err)
	}

	recorder := httptest.NewRecorder()
	url := "/movies"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("expected no error while crearting request, but got %v", err)
	}

	q := request.URL.Query()
	q.Add("page_id", "1")
	q.Add("page_size", "2")
	q.Add("actor", "leo")
	request.URL.RawQuery = q.Encode()

	server.router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code 200 but got %d", recorder.Code)
	}
}

func TestListMoviesShouldReturnMoviesFilteredByGenreAndActor(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStore(ctrl)
	mockMovies := initMockMovies()
	mockReader := mocks.NewMockMoviesReader(ctrl)
	mockCartService := mocks.NewMockICartService(ctrl)
	mockPriceService := mocks.NewMockIPriceService(ctrl)
	mockStore.EXPECT().ListMovies(gomock.Any(), gomock.Any()).Return(mockMovies, nil).Times(1)

	gin.SetMode(gin.TestMode)
	server, err := NewServer(mockStore, mockReader, mockCartService, mockPriceService)
	if err != nil {
		t.Errorf("expected no error while creating server, but got %v", err)
	}

	recorder := httptest.NewRecorder()
	url := "/movies"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("expected no error while crearting request, but got %v", err)
	}

	q := request.URL.Query()
	q.Add("page_id", "1")
	q.Add("page_size", "2")
	q.Add("genre", "drama")
	q.Add("actor", "leo")
	request.URL.RawQuery = q.Encode()

	server.router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code 200 but got %d", recorder.Code)
	}
}

func TestListMoviesShouldReturnMoviesFilteredByGenreActorAndYear(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStore(ctrl)
	mockMovies := initMockMovies()
	mockReader := mocks.NewMockMoviesReader(ctrl)
	mockCartService := mocks.NewMockICartService(ctrl)
	mockPriceService := mocks.NewMockIPriceService(ctrl)
	mockStore.EXPECT().ListMovies(gomock.Any(), gomock.Any()).Return(mockMovies, nil).Times(1)

	gin.SetMode(gin.TestMode)
	server, err := NewServer(mockStore, mockReader, mockCartService, mockPriceService)
	if err != nil {
		t.Errorf("expected no error while creating server, but got %v", err)
	}

	recorder := httptest.NewRecorder()
	url := "/movies"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("expected no error while crearting request, but got %v", err)
	}

	q := request.URL.Query()
	q.Add("page_id", "1")
	q.Add("page_size", "2")
	q.Add("year", "1972")
	q.Add("genre", "drama")
	q.Add("actor", "leo")
	request.URL.RawQuery = q.Encode()

	server.router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code 200 but got %d", recorder.Code)
	}
}

func TestGetMovieShouldReturnMovieDetailsWhenIdIsCorrect(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStore(ctrl)
	mockMovies := initMockMovies()
	id := mockMovies[0].ID
	mockReader := mocks.NewMockMoviesReader(ctrl)
	mockCartService := mocks.NewMockICartService(ctrl)
	mockPriceService := mocks.NewMockIPriceService(ctrl)
	mockStore.EXPECT().GetMovie(gomock.Any(), gomock.Eq(id)).Return(mockMovies[0], nil).Times(1)

	gin.SetMode(gin.TestMode)
	server, err := NewServer(mockStore, mockReader, mockCartService, mockPriceService)
	if err != nil {
		t.Errorf("expected no error while creating server, but got %v", err)
	}

	recorder := httptest.NewRecorder()
	url := fmt.Sprintf("/movies/%s", id)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("expected no error while crearting request, but got %v", err)
	}

	server.router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code 200 but got %d", recorder.Code)
	}
}

func TestGetMovieShouldInternalServerErrorWhenDbCallReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStore(ctrl)
	mockMovies := initMockMovies()
	id := mockMovies[0].ID
	mockReader := mocks.NewMockMoviesReader(ctrl)
	mockCartService := mocks.NewMockICartService(ctrl)
	mockPriceService := mocks.NewMockIPriceService(ctrl)
	mockStore.EXPECT().GetMovie(gomock.Any(), gomock.Eq(id)).Return(db.Movie{}, errors.New("mock db connection error")).Times(1)

	gin.SetMode(gin.TestMode)
	server, err := NewServer(mockStore, mockReader, mockCartService, mockPriceService)
	if err != nil {
		t.Errorf("expected no error while creating server, but got %v", err)
	}

	recorder := httptest.NewRecorder()
	url := fmt.Sprintf("/movies/%s", id)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("expected no error while crearting request, but got %v", err)
	}

	server.router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusInternalServerError {
		t.Errorf("expected status code 500 but got %d", recorder.Code)
	}
}

func TestGetMovieShouldReturnNotFoundWhenMovieIsNotPresentInDatabase(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStore(ctrl)
	mockMovies := initMockMovies()
	id := mockMovies[0].ID
	mockReader := mocks.NewMockMoviesReader(ctrl)
	mockCartService := mocks.NewMockICartService(ctrl)
	mockStore.EXPECT().GetMovie(gomock.Any(), gomock.Eq(id)).Return(db.Movie{}, sql.ErrNoRows).Times(1)
	mockPriceService := mocks.NewMockIPriceService(ctrl)

	gin.SetMode(gin.TestMode)
	server, err := NewServer(mockStore, mockReader, mockCartService, mockPriceService)
	if err != nil {
		t.Errorf("expected no error while creating server, but got %v", err)
	}

	recorder := httptest.NewRecorder()
	url := fmt.Sprintf("/movies/%s", id)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("expected no error while crearting request, but got %v", err)
	}

	server.router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusNotFound {
		t.Errorf("expected status code 404 but got %d", recorder.Code)
	}
}
