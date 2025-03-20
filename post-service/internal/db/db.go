package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Post struct {
	ID      int64
	UserID  int64
	Content string
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
		CREATE TABLE IF NOT EXISTS posts (
			id SERIAL PRIMARY KEY,
			user_id BIGINT NOT NULL,
			content TEXT NOT NULL
		)`

	_, err := db.Exec(query)
	return err
}

func (d *Database) GetPost(id int64) (*Post, error) {
	var post Post
	err := d.db.QueryRow("SELECT id, user_id, content FROM posts WHERE id = $1", id).Scan(&post.ID, &post.UserID, &post.Content)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error getting post: %v", err)
	}
	return &post, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}
