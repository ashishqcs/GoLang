package service

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"movierental/client"
	"movierental/config"
	"movierental/model"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

var ErrMovieDoesNotExist = errors.New("could not find movie")

var instance *MovieService

type Service interface {
	GetAllMovies(ctx *gin.Context)
	GetMovie(ctx *gin.Context, id string)
}

type MovieService struct {
	movies   []model.Movie
	movieMap map[string]*model.Movie
	Client   client.Client
	config   config.Config
}

func NewMovieService(config config.Config) *MovieService {
	var once sync.Once
	once.Do(func() {
		instance = &MovieService{
			Client: client.NewOmdbClient(config),
			config: config,
		}
		instance.loadMoviesFromCsv()
	})
	return instance
}

func (me *MovieService) GetAllMovies(ctx *gin.Context) {
	movieResponse := []model.MovieResponse{}
	for _, movie := range me.movies {
		movieResponse = append(movieResponse, *model.MovieToMovieResponse(movie))
	}
	ctx.JSON(http.StatusOK, movieResponse)
}

func (me *MovieService) GetMovie(ctx *gin.Context, id string) {
	movie, ok := me.movieMap[id]
	if !ok {

		clientMovieResponse, err := me.Client.GetMovieById(id)

		if err != nil || strings.ToLower(clientMovieResponse.Response) == "false" {
			ctx.JSON(http.StatusNotFound, errorResponse(ErrMovieDoesNotExist))
			return
		}

		movie = model.ClientMovieResponseToMovie(*clientMovieResponse)
		me.movies = append(me.movies, *movie)
		me.movieMap[movie.Id] = movie
		go me.saveMovieToCsv(movie)
	}
	ctx.JSON(http.StatusOK, movie)

}

func (me *MovieService) saveMovieToCsv(movie *model.Movie) {
	var mu sync.Mutex
	movieArr := []string{
		movie.Id,
		movie.Title,
		movie.Year,
		movie.Rated,
		movie.Released,
		movie.Genre,
		movie.Actors,
		movie.Language,
		movie.Country,
		movie.Poster,
		movie.Type,
	}

	mu.Lock()
	file, err := os.OpenFile(me.config.MovieFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("Error writing to csv: ", err)
		return
	}
	w := csv.NewWriter(file)
	w.Comma = ';'
	w.Write(movieArr)
	w.Flush()
	mu.Unlock()
}

func (me *MovieService) loadMoviesFromCsv() {
	movies := []model.Movie{}
	movieMap := make(map[string]*model.Movie, 0)

	file, err := os.Open(me.config.MovieFilePath)

	if err != nil {
		log.Fatalf("Unable to load movie data: %s", err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'

	for i := 0; ; i++ {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if i > 0 {
			movie := model.MovieFromString(rec)
			movies = append(movies, *movie)
			movieMap[movie.Id] = movie
		}
	}
	me.movies = movies
	me.movieMap = movieMap
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
