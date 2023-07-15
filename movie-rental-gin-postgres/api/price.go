package api

import (
	"fmt"
	"movieRentals/model"
	"movieRentals/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type pricingRequest struct {
	Items []pricingItemRequest `json:"items" binding:"dive"`
}

type pricingItemRequest struct {
	MovieId  string `json:"id" binding:"required"`
	RentFrom string `json:"from" binding:"date2"`
	RentTo   string `json:"to" binding:"date2"`
}

func (s *Server) findPrice(ctx *gin.Context) {
	var req pricingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	items := mapPricingReqToPricingItem(req)

	response, err := s.priceService.PriceItem(ctx, items)

	if err == service.ErrInvalidDates {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	fmt.Printf("pricing request %v", req)
	ctx.JSON(http.StatusOK, response)
}

func mapPricingReqToPricingItem(req pricingRequest) []model.PricingItem {
	items := []model.PricingItem{}
	for _, item := range req.Items {
		priceItem := model.PricingItem{
			MovieId:  item.MovieId,
			RentFrom: item.RentFrom,
			RentTo:   item.RentTo,
		}
		items = append(items, priceItem)
	}

	return items
}
