package api

import (
	"log"
	"movieRentals/api/validators"
	db "movieRentals/db/sqlc"
	"movieRentals/model"
	"movieRentals/service"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type MoviesReader interface {
	GetMovies() (*model.Movies, error)
}

type Server struct {
	reader       MoviesReader
	store        db.Store
	cartService  service.ICartService
	priceService service.IPriceService
	router       *gin.Engine
}

func NewServer(db db.Store, reader MoviesReader, cartService service.ICartService, priceService service.IPriceService) (*Server, error) {
	server := &Server{
		store:        db,
		cartService:  cartService,
		priceService: priceService,
		reader:       reader,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("year", validators.ValidYear)
		v.RegisterValidation("date2", validators.ValidDate2)
	}

	server.setupRouter()
	return server, nil
}

func (s *Server) setupRouter() {
	router := gin.Default()
	router.GET("/movies", s.listMovies)
	router.GET("/movies/:id", s.getMovie)
	router.POST("/cart/movies", s.addToCart)
	router.GET("/cart/movies", s.getCart)
	router.POST("/price", s.findPrice)
	s.router = router
}

func (s *Server) Start(address string) error {
	err := s.load()
	if err != nil {
		log.Fatalf("unable to load seed data: %v", err)
	}
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
