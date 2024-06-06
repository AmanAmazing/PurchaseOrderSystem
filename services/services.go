package services

import (
	"PurchaseOrderSystem/utils"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func InitDB() error {
	var err error
	dsn := fmt.Sprintf("%s://%s:%s@localhost:%s/%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		return err
	}
	return db.Ping(context.Background())
}

func TestData() {
	err := utils.TestDB(db)
	if err != nil {
		log.Fatalf("Error occurred in the utils.TestDB() function: %v", err)
	}
}
