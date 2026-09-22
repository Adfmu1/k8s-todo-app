package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"
	"todo-backend/postgres"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type app struct {
	server  *http.Server
	logger  *slog.Logger
	dbQuery *postgres.Queries
}

var application app

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	application.logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))

	// DB block
	ctx := context.Background()
	dbConn, err := pgx.Connect(ctx, os.Getenv("POSTGRES_URL"))
	checkErr(err, *application.logger)
	defer dbConn.Close(ctx)

	application.dbQuery = postgres.New(dbConn)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /todos", application.getTodos)
	mux.HandleFunc("POST /todos", application.postTodos)

	serv := &http.Server{
		Addr:        ":" + port,
		Handler:     mux,
		ReadTimeout: 5 * time.Second,
	}

	application.logger.Info("Serving backend", "PORT", port)
	serv.ListenAndServe()
}
