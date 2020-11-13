package pgdb

import (
	"fmt"
	"log"

	_ "github.com/lib/pq" //postgres driver

	"github.com/jmoiron/sqlx"
)

// DB ...
// var DB *sqlx.DB

// GetDB ...
func GetDB() *sqlx.DB {
	var DB *sqlx.DB
	if DB == nil {
		dbPath := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			"localhost", 5432, "postgres", "postgres", "postgres")
		DB, err := sqlx.Open("postgres", dbPath)
		if err != nil {
			log.Fatal(err.Error())
		}
		if err = DB.Ping(); err != nil {
			log.Fatal(err.Error())
		}
		return DB
	}
	return DB
}
