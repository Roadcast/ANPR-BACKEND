package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"entgo.io/ent/dialect/sql/schema"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/vektah/gqlparser/v2/ast"
	"go-ent-project/config"
	"go-ent-project/graph"
	"go-ent-project/internal/ent"
	"go-ent-project/utils/celery"
	"go-ent-project/utils/db"
	"go-ent-project/utils/middleware"
	redisDB "go-ent-project/utils/redis"
	"log"
	"net/http"
	"os"
)

import _ "go-ent-project/internal/ent/runtime"

const defaultPort = "8000"

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, loading from environment variables.")
	}

	cfg := config.LoadDBConfig()
	dbURL := config.GetPostgresDSN(cfg)

	client, err := ent.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	err = client.Schema.Create(context.Background(), schema.WithGlobalUniqueID(true))
	if err != nil {
		log.Fatalf("failed to create schema resources: %v", err)
	}

	redisDB.Initialize()
	println("dbURL: ", dbURL)
	// Run migrations
	//db.RunMigrations(dbURL, "./migrations")

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			Client: db.Client, // Pass the initialized Ent client
		},
	}))
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.HandleFunc("/car", carHandler)
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", middleware.AuthMiddleware(client)(srv))

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	celery.Initialize()
	defer celery.ShutDown()
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler(http.DefaultServeMux)))

}

func applyUUIDDefaults(db *sql.DB, tables ...string) error {
	_, err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
	if err != nil {
		return fmt.Errorf("failed to enable uuid-ossp extension: %w", err)
	}

	for _, table := range tables {
		query := fmt.Sprintf(`ALTER TABLE %s ALTER COLUMN id SET DEFAULT uuid_generate_v4()`, table)
		if _, err := db.Exec(query); err != nil {
			return fmt.Errorf("failed to set default on table %s: %w", table, err)
		}
	}
	return nil
}

func carHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var body struct {
		PlateNumber string `json:"plate_number"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Received alert for plate: %s", body.PlateNumber)

	// You can now log, insert into DB, trigger downstream alerts, etc.
	// For now, respond with acknowledgment
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Received plate: " + body.PlateNumber))
}
