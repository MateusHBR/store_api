package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PQDatabase struct{}

func (pq PQDatabase) OpenConnection() (*sql.DB, error) {
	// dbUser := os.Getenv("POSTGRES_USER")
	// dbPasswd := os.Getenv("POSTGRES_PASSWORD")
	// dbName := os.Getenv("POSTGRES_DB")

	connStr := "host=localhost port=5432 user=mallus password=mysecretpassword dbname=mallus_db sslmode=disable"

	return sql.Open("postgres", connStr)
}
