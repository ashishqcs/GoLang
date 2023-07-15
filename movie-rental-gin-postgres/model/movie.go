package model

type Movies struct {
	Movies []Movie `json:"movies"`
}

type Movie struct {
	ID       string      `json:"id"`
	Title    string      `json:"title"`
	Released ReleaseDate `json:"released"`
	Genre    string      `json:"genre"`
	Actors   string      `json:"actors"`
	Year     int32       `json:"year"`
	Price    int64       `json:"price"`
	Quantity int32       `json:"quantity"`
}
