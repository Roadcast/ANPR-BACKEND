package db

import (
	"database/sql"
	"go-ent-project/config"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgresConnection(cfg *config.DBConfig) *sql.DB {
	dsn := config.GetPostgresDSN(cfg)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Successfully connected to the database.")
	return db
}
