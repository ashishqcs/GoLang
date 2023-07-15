package reader

import (
	"strings"
	"testing"
	"time"
)

func TestGetMoviesFromReaderReturnsErrorWhenInputIsNotInCorrectFormat(t *testing.T) {
	mockReader := strings.NewReader("test")
	movies, err := getMoviesFromReader(mockReader)
	if movies != nil || err == nil {
		t.Errorf("expected error but got none")
	}
}

func TestGetMoviesFromReaderReturnsMoviesWhenInputIsInCorrectFormat(t *testing.T) {
	mockJson := "{\"movies\":[{\"id\": \"tt0068646\",\"title\": \"The Godfather\",\"released\": \"24 Mar 1972\",\"genre\": \"Crime, Drama\",\"actors\": \"Marlon Brando, Al Pacino, James Caan\",\"year\": 1972,\"price\": 20,\"quantity\": 10}]}"
	mockReader := strings.NewReader(mockJson)
	movies, err := getMoviesFromReader(mockReader)
	if err != nil || movies == nil || len(movies.Movies) != 1 {
		t.Errorf("expected movies with length %d but got %d", 1, len(movies.Movies))
	}
	if movies.Movies[0].ID != "tt0068646" {
		t.Errorf("expected movie with id %s but got %s", "tt0068646", movies.Movies[0].ID)
	}
	rd, _ := time.Parse("02 Jan 2006", "24 Mar 1972")

	if movies.Movies[0].Released.Time != rd {
		t.Errorf("expected movie with release date %s but got %s", rd, movies.Movies[0].Released.Time)
	}
}
