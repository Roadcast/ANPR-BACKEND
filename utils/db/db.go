package db

import (
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	"go-ent-project/config"
	"go-ent-project/internal/ent"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"log"
)

var Client *ent.Client

func init() {
	var err error
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, loading from environment variables.")
	}

	cfg := config.LoadDBConfig()
	dbURL := config.GetPostgresDSN(cfg)

	db, err := otelsql.Open(
		"postgres",
		dbURL,
		otelsql.WithAttributes(semconv.DBSystemPostgreSQL),
	)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	_, err = db.Query("SELECT 1")
	if err != nil {
		log.Fatalf("failed to ping database: %v", err)
		return
	}

	Client = ent.NewClient(ent.Driver(sql.OpenDB(dialect.Postgres, db)))
}
