package postgres

import (
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var ErrFailedToInsertData = errors.New("failed to insert data")
var ErrFailedToFetchData = errors.New("failed to get data")

type Repository interface {
	SaveBill(bill *BillEntity) error
	SaveBills(bills []BillEntity) error
	GetBillsByName(name string) ([]BillEntity, error)
}

func NewBillRepository() BillRepository {
	db, _ := getDB()
	return BillRepository{
		DB: db,
	}
}

type BillRepository struct {
	*sqlx.DB
}

func (b BillRepository) SaveBill(bill *BillEntity) error {
	if _, err := b.NamedQuery("INSERT INTO bill (id, name, address, plan_name, date, amount) VALUES(:id, :name, :address, :plan_name, :date, :amount)", bill); err != nil {
		log.Printf("unable to save data for id %d. Err: %s", bill.Id, err.Error())
		return ErrFailedToInsertData
	}

	return nil
}

func (b BillRepository) SaveBills(bills []BillEntity) error {
	if _, err := b.NamedExec("INSERT INTO bill (id, name, address, plan_name, date, amount) "+
		"VALUES(:id, :name, :address, :plan_name, :date, :amount)", bills); err != nil {
		log.Printf("unable to save data. Err: %s", err.Error())
		return ErrFailedToInsertData
	}

	return nil
}

func (b BillRepository) GetBillsByName(name string) ([]BillEntity, error) {
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
