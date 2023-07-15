package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"movierental/config"
	"movierental/model"
	"net/http"
	"time"
)

var ErrServerNotResponding = errors.New("server not responding")
var ErrFoundNoResponse = errors.New("no Response")
var ErrParsingData = errors.New("could not parse response")

type Client interface {
	GetMovieById(id string) (*model.ClientMovieResponse, error)
}

type OmdbClient struct {
	Client http.Client
	config config.Config
}

func NewOmdbClient(config config.Config) *OmdbClient {
	return &OmdbClient{
		config: config,
	}
}

func (ip *OmdbClient) GetMovieById(id string) (*model.ClientMovieResponse, error) {

	url := fmt.Sprintf("%s/?apikey=%s&i=%s", ip.config.OmdbClientUrl, ip.config.OmdbApiKey, id)
	start := time.Now()
	res, err := ip.Client.Get(url)
	log.Printf("client took: %v", time.Since(start))

	if err != nil {
		log.Printf("Error: %s", err.Error())
		return nil, ErrServerNotResponding
	}

	body, _ := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		log.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return nil, ErrFoundNoResponse
	}

	responseMap := make(map[string]interface{})

	start = time.Now()
	if err = json.Unmarshal(body, &responseMap); err != nil {
		log.Printf("Could not parse response body: %s", err.Error())
		return nil, ErrParsingData
	}
	log.Printf("response parsing took: %v", time.Since(start))

	return mapBodyToClientMovieResponse(responseMap), nil
}

func mapBodyToClientMovieResponse(response map[string]interface{}) *model.ClientMovieResponse {
	return &model.ClientMovieResponse{
		Title:      GetOrDefault(response, "Title", ""),
		Year:       GetOrDefault(response, "Year", ""),
		Rated:      GetOrDefault(response, "Rated", ""),
		Released:   GetOrDefault(response, "Released", ""),
		Runtime:    GetOrDefault(response, "Runtime", ""),
		Genre:      GetOrDefault(response, "Genre", ""),
		Director:   GetOrDefault(response, "Director", ""),
		Writer:     GetOrDefault(response, "Writer", ""),
		Actors:     GetOrDefault(response, "Actors", ""),
		Plot:       GetOrDefault(response, "Plot", ""),
		Language:   GetOrDefault(response, "Language", ""),
		Country:    GetOrDefault(response, "Country", ""),
		Awards:     GetOrDefault(response, "Awards", ""),
		Poster:     GetOrDefault(response, "Poster", ""),
		Metascore:  GetOrDefault(response, "Metascore", ""),
		ImdbRating: GetOrDefault(response, "imdbRating", ""),
		ImdbVotes:  GetOrDefault(response, "imdbVotes", ""),
		ImdbID:     GetOrDefault(response, "imdbID", ""),
		Type:       GetOrDefault(response, "Type", ""),
		DVD:        GetOrDefault(response, "DVD", ""),
		BoxOffice:  GetOrDefault(response, "BoxOffice", ""),
		Production: GetOrDefault(response, "Production", ""),
		Website:    GetOrDefault(response, "Website", ""),
		Response:   GetOrDefault(response, "Response", ""),
	}

}

func GetOrDefault(m map[string]interface{}, key string, def string) string {
	a := m[key]
	if a != nil {
		return a.(string)
	}
	return def
}
