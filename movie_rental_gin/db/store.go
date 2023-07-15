package db

import (
	"errors"
	"log"
	"movierental/model"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var ErrFailedToInsertData = errors.New("failed to insert data")
var ErrFailedToFetchData = errors.New("failed to get data")

type Repository interface {
	SaveMovie(movie *model.MovieEntity) error
	SaveMovies(movies []model.MovieEntity) error
	GetAllMovies() ([]model.Movie, error)
	GetMovieById(id string) (*model.Movie, error)
}

func NewMovieRepository() MovieRepository {
	db, _ := getDB()
	return MovieRepository{
		DB: db,
	}
}

type MovieRepository struct {
	*sqlx.DB
}

func (b *MovieRepository) SaveMovie(movie *model.MovieEntity) error {
	if _, err := b.NamedQuery("INSERT INTO movie (id, name, address, plan_name, date, amount) VALUES(:id, :name, :address, :plan_name, :date, :amount)", bill); err != nil {
		log.Printf("unable to save data for id %d. Err: %s", movie.Id, err.Error())
		return ErrFailedToInsertData
	}

	return nil
}

func (b *MovieRepository) SaveMovies(movies []model.MovieEntity) error {
	if _, err := b.NamedExec("INSERT INTO bill (id, name, address, plan_name, date, amount) "+
		"VALUES(:id, :name, :address, :plan_name, :date, :amount)", movies); err != nil {
		log.Printf("unable to save data. Err: %s", err.Error())
		return ErrFailedToInsertData
	}

	return nil
}

func (b *MovieRepository) GetAllMovies() ([]model.Movie, error) {
	bills := []BillEntity{}
	if err := b.Select(&bills, "select * from bill where name = $1", name); err != nil {
		log.Printf("unable to save data. Err: %s", err.Error())
		return nil, ErrFailedToFetchData
	}

	return bills, nil
}

func (b *MovieRepository) GetMovieById(id string) (*model.Movie, error) {
	bills := []BillEntity{}
	if err := b.Select(&bills, "select * from bill where name = $1", name); err != nil {
		log.Printf("unable to save data. Err: %s", err.Error())
		return nil, ErrFailedToFetchData
	}

	return bills, nil
}

func getDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", "user=user password=password dbname=internet_bills sslmode=disable")

	if err != nil {
		log.Fatalf("Error Opeaning DB Connection: %s", err.Error())
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Cannot connect to DB. Error: %s", err.Error())
		return nil, err
	}

	return db, nil
}
