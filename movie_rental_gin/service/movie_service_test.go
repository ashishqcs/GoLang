package service

import (
	"movierental/config"
	mock "movierental/mocks"
	"movierental/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var testAllMovieResponse = "[{\"Id\":\"tt0848228\",\"Title\":\"The Avengers\",\"Year\":\"2012\",\"Poster\":\"https://m.media-amazon.com/images/M/MV5BNDYxNjQyMjAtNTdiOS00NGYwLWFmNTAtNThmYjU5ZGI2YTI1XkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_SX300.jpg\",\"Type\":\"movie\"},{\"Id\":\"tt4154796\",\"Title\":\"Avengers: Endgame\",\"Year\":\"2019\",\"Poster\":\"https://m.media-amazon.com/images/M/MV5BMTc5MDE2ODcwNV5BMl5BanBnXkFtZTgwMzI2NzQ2NzM@._V1_SX300.jpg\",\"Type\":\"movie\"},{\"Id\":\"tt4154756\",\"Title\":\"Avengers: Infinity War\",\"Year\":\"2018\",\"Poster\":\"https://m.media-amazon.com/images/M/MV5BMjMxNjY2MDU1OV5BMl5BanBnXkFtZTgwNzY1MTUwNTM@._V1_SX300.jpg\",\"Type\":\"movie\"},{\"Id\":\"tt2395427\",\"Title\":\"Avengers: Age of Ultron\",\"Year\":\"2015\",\"Poster\":\"https://m.media-amazon.com/images/M/MV5BMTM4OGJmNWMtOTM4Ni00NTE3LTg3MDItZmQxYjc4N2JhNmUxXkEyXkFqcGdeQXVyNTgzMDMzMTg@._V1_SX300.jpg\",\"Type\":\"movie\"},{\"Id\":\"tt2258647\",\"Title\":\"The Dark Knight\",\"Year\":\"2011\",\"Poster\":\"https://m.media-amazon.com/images/M/MV5BMzIxNzU4NjkwMV5BMl5BanBnXkFtZTgwNDU4NjM4MDE@._V1_SX300.jpg\",\"Type\":\"movie\"}]"
var testMovieDetailByIdResponse = "{\"Id\":\"tt4154756\",\"Title\":\"Avengers: Infinity War\",\"Year\":\"2018\",\"Rated\":\"PG-13\",\"Released\":\"27 Apr 2018\",\"Genre\":\"Action, Adventure, Sci-Fi\",\"Actors\":\"Robert Downey Jr., Chris Hemsworth, Mark Ruffalo\",\"Language\":\"English\",\"Country\":\"United States\",\"Poster\":\"https://m.media-amazon.com/images/M/MV5BMjMxNjY2MDU1OV5BMl5BanBnXkFtZTgwNzY1MTUwNTM@._V1_SX300.jpg\",\"Type\":\"movie\"}"

func TestShouldLoadDataFromFile(t *testing.T) {
	service := &MovieService{
		config: getTestConfig(),
		Client: new(mock.Client),
	}

	service.loadMoviesFromCsv()

	assert.Equal(t, 5, len(service.movies))
	assertAllMovieFields(t, service.movies[0])
}

func TestShouldGetAllMovies(t *testing.T) {
	service := NewMovieService(getTestConfig())

	ctx, w := getTestGinContext()

	service.GetAllMovies(ctx)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Equal(t, testAllMovieResponse, w.Body.String())
}

func TestShouldGetMovieDetailsById(t *testing.T) {
	service := NewMovieService(getTestConfig())

	ctx, w := getTestGinContext()

	service.GetMovie(ctx, "tt4154756")

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Equal(t, testMovieDetailByIdResponse, w.Body.String())
}

// func TestShouldGetMovieFromOmdbIfNotPresent(t *testing.T) {
// 	service := NewMovieService(getTestConfig())
// 	service.Client = new(mock.Client)

// 	ctx, w := getTestGinContext()

// 	service.GetMovie(ctx, "tt4154756")

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	assert.Equal(t, testMovieDetailByIdResponse, w.Body.String())
// }

func getTestGinContext() (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	return ctx, w
}

func getTestConfig() config.Config {
	return config.Config{
		MovieFilePath: "/Users/ashish.singh/bootcamp/movie-rental-service/resources/movies_test.csv",
	}
}

func assertAllMovieFields(t *testing.T, movie model.Movie) {
	assert.Equal(t, "tt0848228", movie.Id)
	assert.Equal(t, "The Avengers", movie.Title)
	assert.Equal(t, "2012", movie.Year)
	assert.Equal(t, "PG-13", movie.Rated)
	assert.Equal(t, "04 May 2012", movie.Released)
	assert.Equal(t, "Action, Sci-Fi", movie.Genre)
	assert.Equal(t, "Robert Downey Jr., Chris Evans, Scarlett Johansson", movie.Actors)
	assert.Equal(t, "English, Russian", movie.Language)
	assert.Equal(t, "United States", movie.Country)
	assert.Equal(t, "https://m.media-amazon.com/images/M/MV5BNDYxNjQyMjAtNTdiOS00NGYwLWFmNTAtNThmYjU5ZGI2YTI1XkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_SX300.jpg", movie.Poster)
	assert.Equal(t, "movie", movie.Type)
}
