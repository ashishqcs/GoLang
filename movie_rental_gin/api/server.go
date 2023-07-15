package api

import (
	"movierental/service"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router       *gin.Engine
	movieService service.Service
}

func NewServer(movieService service.Service) *Server {
	server := &Server{
		movieService: movieService,
	}

	router := gin.Default()

	router.GET("/movies", server.getMovies)
	router.GET("/movie/:id", server.getMovie)

	server.router = router
	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
