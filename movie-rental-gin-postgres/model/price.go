package model

type PricingResponse struct {
	Movies     []PricingItemResponse
	TotalPrice int64
}

type PricingItemResponse struct {
	MovieId      string
	MoviePrice   int64
	RentDuration string
}

type PricingItem struct {
	MovieId  string
	RentFrom string
	RentTo   string
}

func NewPricingResponse(movies []PricingItemResponse, price int64) *PricingResponse {
	return &PricingResponse{
		Movies:     movies,
		TotalPrice: price,
	}
}

func NewPricingItemResponse(movieId string, rentPrice int64, rentDuration string) *PricingItemResponse {
	return &PricingItemResponse{
		MovieId:      movieId,
		MoviePrice:   rentPrice,
		RentDuration: rentDuration,
	}
}
