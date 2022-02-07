package database

import (
	"database/sql"
)

type DatabaseAdapter interface {
	OpenConnection() (*sql.DB, error)
}
