package client

import (
	"movierental/config"
	"movierental/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	movieDetail := model.ClientMovieResponse{
		ImdbID: "tt1234",
		Title:  "Thor",
	}

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.URL.String(), "/?apikey=testKey&i=tt1234")
		response := []byte("{\"imdbID\":\"tt1234\",\"Title\":\"Thor\"}")
		rw.Write(response)
	}))

	defer server.Close()

	client := OmdbClient{
		Client: *server.Client(),
		config: config.Config{OmdbClientUrl: server.URL, OmdbApiKey: "testKey"},
	}

	response, _ := client.GetMovieById("tt1234")

	assert.Equal(t, movieDetail.ImdbID, response.ImdbID)
	assert.Equal(t, movieDetail.Title, response.Title)
}
