package postgres

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func TestDataBase(t *testing.T) {
	bill := BillEntity{
		Id:       1,
		Name:     "Ram",
		Address:  "ABSD",
		Date:     "12",
		Amount:   22.5,
		PlanName: "Fiber",
	}

	db, _ := sqlx.Open("postgres", "user=user password=password dbname=internet_bills sslmode=disable")
	repo := BillRepository{
		DB: db,
	}

	repo.SaveBill(&bill)
}

func TestBatchSave(t *testing.T) {
	bills := []BillEntity{
		{
			Id:       1,
			Name:     "Ram",
			Address:  "ABSD",
			Date:     "12",
			Amount:   22.5,
			PlanName: "Fiber",
		},
		{
			Id:       2,
			Name:     "Manas",
			Address:  "ABSD",
			Date:     "12",
			Amount:   22.5,
			PlanName: "Fiber",
		},
	}

	db, _ := sqlx.Open("postgres", "user=user password=password dbname=internet_bills sslmode=disable")
	repo := BillRepository{
		DB: db,
	}

	repo.SaveBills(bills)
}

func TestFetchByName(t *testing.T) {

	db, _ := sqlx.Open("postgres", "user=user password=password dbname=internet_bills sslmode=disable")
	repo := BillRepository{
		DB: db,
	}

	be, _ := repo.GetBillsByName("Ram")
	t.Log(be)

}
