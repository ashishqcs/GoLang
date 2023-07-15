package api

import (
	"database/sql"
	"errors"
	"movieRentals/api/dto"
	db "movieRentals/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type listMoviesRequest struct {
	PageID   int32  `form:"page_id" binding:"required,min=1"`
	PageSize int32  `form:"page_size" binding:"required,min=1,max=5"`
	Genre    string `form:"genre"`
	Actor    string `form:"actor"`
	Year     int    `form:"year" binding:"year"`
}

func (s *Server) listMovies(ctx *gin.Context) {
	var req listMoviesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListMoviesParams{
		Limit:    req.PageSize,
		Offset:   (req.PageID - 1) * req.PageSize,
		LkGenre:  req.Genre != "",
		Genre:    sql.NullString{String: req.Genre, Valid: req.Genre != ""},
		LkActors: req.Actor != "",
		Actor:    sql.NullString{String: req.Actor, Valid: req.Actor != ""},
		EqYear:   req.Year != 0,
		Year:     int32(req.Year),
	}

	movies, err := s.store.ListMovies(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, dto.MapMoviesFromDbModel(movies))
}

type getMovieRequest struct {
	Id string `uri:"id" binding:"required"`
}

func (s *Server) getMovie(ctx *gin.Context) {
	var req getMovieRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	movie, err := s.store.GetMovie(ctx, req.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, dto.MapMovieFromDbModel(&movie))
}
