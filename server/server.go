package server

import "database/sql"

type Server struct {
	DB *sql.DB
}

func New() Server {
	return Server{}
}

func (s Server) WithDatabaseConnection(db *sql.DB) Server {
	s.DB = db

	return s
}
