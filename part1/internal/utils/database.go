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

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	conn.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	// Create user table
	_, err = conn.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			"id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			"name" text NOT NULL,
			email text NOT NULL,
			age INT NOT NULL
		);
	`)
	if err != nil {
		return nil, err
	}

	return &DB{Conn: conn}, nil
}

func (db *DB) Close() {
	db.Conn.Close()
}
