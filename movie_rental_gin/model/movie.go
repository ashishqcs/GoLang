package model

type Movie struct {
	Id       string
	Title    string
	Year     string
	Rated    string
	Released string
	Genre    string
	Actors   string
	Language string
	Country  string
	Poster   string
	Type     string
}

type MovieEntity struct {
	Id       string `db:"id"`
	Title    string
	Year     string
	Rated    string
	Released string
	Genre    string
	Actors   string
	Language string
	Country  string
	Poster   string
	Type     string
}

type MovieResponse struct {
	Id     string
	Title  string
	Year   string
	Poster string
	Type   string
}

type MovieDetailResponse struct {
	Id       string
	Title    string
	Year     string
	Rated    string
	Released string
	Genre    string
	Actors   string
	Language string
	Country  string
	Poster   string
	Type     string
}

type ClientMovieRating struct {
	Source string
	Value  string
}

type ClientMovieResponse struct {
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Poster     string
	Ratings    []ClientMovieRating
	Metascore  string
	ImdbRating string
	ImdbVotes  string
	ImdbID     string
	Type       string
	DVD        string
	BoxOffice  string
	Production string
	Website    string
	Response   string
}

func MovieFromString(splitMovie []string) *Movie {
	// splitMovie := strings.Split(s[0], ";")

	return &Movie{
		Id:       splitMovie[0],
		Title:    splitMovie[1],
		Year:     splitMovie[2],
		Rated:    splitMovie[3],
		Released: splitMovie[4],
		Genre:    splitMovie[5],
		Actors:   splitMovie[6],
		Language: splitMovie[7],
		Country:  splitMovie[8],
		Poster:   splitMovie[9],
		Type:     splitMovie[10],
	}
}

func MovieToMovieDetailResponse(m Movie) *MovieDetailResponse {

	return &MovieDetailResponse{
		Id:       m.Id,
		Title:    m.Title,
		Year:     m.Year,
		Rated:    m.Rated,
		Released: m.Released,
		Genre:    m.Genre,
		Actors:   m.Actors,
		Language: m.Language,
		Country:  m.Country,
		Poster:   m.Poster,
		Type:     m.Type,
	}
}

func MovieToMovieResponse(movie Movie) *MovieResponse {
	return &MovieResponse{
		Id:     movie.Id,
		Title:  movie.Title,
		Year:   movie.Year,
		Poster: movie.Poster,
		Type:   movie.Type,
	}
}

func ClientMovieResponseToMovie(m ClientMovieResponse) *Movie {
	return &Movie{
		Id:       m.ImdbID,
		Title:    m.Title,
		Year:     m.Year,
		Rated:    m.Rated,
		Released: m.Released,
		Genre:    m.Genre,
		Actors:   m.Actors,
		Language: m.Language,
		Country:  m.Country,
		Poster:   m.Poster,
		Type:     m.Type,
	}
}
