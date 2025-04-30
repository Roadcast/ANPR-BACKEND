package main

import (
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
	//postgresDB := db.NewPostgresConnection(cfg)
	client, err := ent.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	//err = client.Schema.Create(context.Background(), schema.WithGlobalUniqueID(true))
	//if err != nil {
	//	log.Fatalf("failed to create schema resources: %v", err)
	//}
	//defer func(postgresDB *sql.DB) {
	//	err := postgresDB.Close()
	//	if err != nil {
	//		log.Fatalf("failed to close database: %v", err)
	//	}
	//}(postgresDB)
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
