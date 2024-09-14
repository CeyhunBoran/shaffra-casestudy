package utils

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type DB struct {
	Conn *sql.DB
}

func NewDB(connectionString string) (*DB, error) {
	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(25)
	conn.SetMaxIdleConns(25)
	conn.SetConnMaxLifetime(time.Hour)
	conn.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return &DB{Conn: conn}, nil
}

func (db *DB) Close() {
	db.Conn.Close()
}
