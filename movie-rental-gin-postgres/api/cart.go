package api

import (
	"errors"
	"movieRentals/model"
	"movieRentals/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type addToCartRequest struct {
	MovieIds []string `json:"movies"`
}

func (s *Server) addToCart(ctx *gin.Context) {
	var req addToCartRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := s.cartService.AddToCart(ctx, req.MovieIds)

	if errors.Is(err, model.ErrInvalidMovieId) ||
		errors.Is(err, model.ErrMovieIdNotFound) ||
		errors.Is(err, service.ErrNoMovieIdsPresent) {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, req.MovieIds)
}

func (s *Server) getCart(ctx *gin.Context) {
	cart := s.cartService.GetCart(ctx)

	ctx.JSON(http.StatusOK, *cart)
}
