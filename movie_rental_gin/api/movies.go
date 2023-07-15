package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ErrInvalidRequest = errors.New("not a valid request")

type MovieIdRequest struct {
	Id string `uri:"id" binding:"required"`
}

func (s *Server) getMovies(ctx *gin.Context) {
	s.movieService.GetAllMovies(ctx)
}

func (s *Server) getMovie(ctx *gin.Context) {
	var request MovieIdRequest
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(ErrInvalidRequest))
		return
	}
	s.movieService.GetMovie(ctx, request.Id)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
