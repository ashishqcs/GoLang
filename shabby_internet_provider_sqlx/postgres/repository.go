package postgres

// import (
// 	"log"
// 	"sync"

// 	"github.com/jmoiron/sqlx"
// 	_ "github.com/lib/pq"
// )

// var once sync.Once
// var billRepository BillRepository

// func GetBillRepository() *BillRepository {
// 	once.Do(func() {
// 		db, _ := GetDB()
// 		billRepository = *NewBillRepository(db)
// 	})

// 	return &billRepository
// }

// func GetDB() (*sqlx.DB, error) {
// 	db, err := sqlx.Open("postgres", "user=user password=password dbname=internet_bills sslmode=disable")

// 	if err != nil {
// 		log.Fatalf("Error Opeaning DB Connection: %s", err.Error())
// 		return nil, err
// 	}

// 	if err = db.Ping(); err != nil {
// 		log.Fatalf("Cannot connect to DB. Error: %s", err.Error())
// 		return nil, err
// 	}

// 	return db, nil
// }

// type Repository struct {
// 	Bill
// }
