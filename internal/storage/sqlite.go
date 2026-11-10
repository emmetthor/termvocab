package storage

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Word struct {
	ID int
	Word string
	Meaning string
}

type Storage struct {
	db *sql.DB
}

// Open_db opens a SQL database and initializes it.
//
// The database file is created if it does not exist.
func Open_db(path string) (*Storage, error) {
	db, err := sql.Open("sqlite", path)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	query := `
	CREATE TABLE IF NOT EXISTS vocabulary (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		word TEXT NOT NULL UNIQUE COLLATE NOCASE,
		meaning TEXT NOT NULL
	);`

	_, err = db.Exec(query)

	if err != nil {
		db.Close()
		return nil, fmt.Errorf("create table: %w", err)
	}

	return &Storage{db: db}, nil
}