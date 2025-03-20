package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	ID   int64
	Name string
}

type Database struct {
	db *sql.DB
}

func NewDatabase(dsn string) (*Database, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	if err := createTable(db); err != nil {
		return nil, fmt.Errorf("error creating table: %v", err)
	}

	return &Database{db: db}, nil
}

func createTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL
		)`

	_, err := db.Exec(query)
	return err
}

func (d *Database) GetUser(id int64) (*User, error) {
	var user User
	err := d.db.QueryRow("SELECT id, name FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error getting user: %v", err)
	}
	return &user, nil
}

func (d *Database) CreateUser(name string) (*User, error) {
	var user User
	query := `INSERT INTO users (name) VALUES ($1) RETURNING id, name`

	err := d.db.QueryRow(query, name).Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %v", err)
	}

	return &user, nil
}

func (d *Database) ListUsers(pageSize int32, page int32) ([]*User, int32, error) {
	if pageSize <= 0 {
		pageSize = 10
	}
	if page <= 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	var totalCount int32
	countErr := d.db.QueryRow("SELECT COUNT(*) FROM users").Scan(&totalCount)
	if countErr != nil {
		return nil, 0, fmt.Errorf("error counting users: %v", countErr)
	}

	query := `SELECT id, name FROM users ORDER BY id LIMIT $1 OFFSET $2`
	rows, err := d.db.Query(query, pageSize, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("error listing users: %v", err)
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, 0, fmt.Errorf("error scanning user row: %v", err)
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("error iterating user rows: %v", err)
	}

	return users, totalCount, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}
