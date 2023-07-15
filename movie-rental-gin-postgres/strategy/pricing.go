package strategy

type PricingStrategy interface {
	PriceItem(days int64) int64
}

type ClassicMoviePricingStrategy struct {
}

type NewMoviePricingStrategy struct {
}

func NewClassicMoviePricingStrategy() *ClassicMoviePricingStrategy {
	return &ClassicMoviePricingStrategy{}
}
func NewNewMoviePricingStrategy() *NewMoviePricingStrategy {
	return &NewMoviePricingStrategy{}
}

func (cmps *NewMoviePricingStrategy) PriceItem(days int64) int64 {
	return 3 + ((days - 2) * 2)
}

func (cmps *ClassicMoviePricingStrategy) PriceItem(days int64) int64 {
	var price int64
	if days > 7 {
		price = (days - 7) + 1
	} else {
		price = 1
	}
	return price
}
