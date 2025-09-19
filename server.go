package main

import (
	"GraphQL/graph"
	"GraphQL/logger"
	"GraphQL/metrics"
	handlerM "GraphQL/src/core/handler"
	"GraphQL/src/core/middleware"
	"GraphQL/src/core/repository"
	"GraphQL/src/core/service"
	pgxhelper "GraphQL/src/pkg"
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vektah/gqlparser/v2/ast"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := pgxhelper.NewPostgresDB(pgxhelper.Config{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("SSLMODE"),
	})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	red, err := pgxhelper.NewRedisDB(context.Background())
	if err != nil {
		log.Fatalf("Error connecting to redis: %v", err)
	}

	repos := repository.NewRepository(db)
	cacheRepo := repository.NewCachedRepo(repos, red, time.Minute)
	services := service.NewService(*cacheRepo, *cacheRepo)

	if err := logger.Init(); err != nil {
		log.Fatalf("failed to init logger: %v", err)
	}
	defer logger.Sync()

	metrics.RegisterMetrics()

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Service: services,
	}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	srv.AroundFields(middlewareMs.GraphQLFieldMetrics)
	srv.AroundResponses(middlewareMs.GraphQLResponseMetrics)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", handlerM.HTTPMetrics(srv))
	http.Handle("/metrics", promhttp.Handler())

	logger.RecordMetric("http_requests_total", 1, map[string]string{
		"path":   "/query",
		"method": "POST",
	})
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
