package main

import (
	"context"
	"log/slog"
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql" // Driver MySQL
	"github.com/Didul-arch/ecom-go-belajar/internal/env"
)

func main() {
	ctx := context.Background()
	cfg := config {
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "user:password_kamu@tcp(127.0.0.1:3306)/ecom?parseTime=true"),
		},
	}

	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))	
	slog.SetDefault(logger)

	// Database
	conn, err := sql.Open("mysql", cfg.db.dsn)
	if err != nil {
		panic(err)
	}

	if err := conn.PingContext(ctx); err != nil {
		panic(err)
	}

	defer conn.Close()

	logger.Info("connected to database", "dsn", cfg.db.dsn)

	api := application{
		config: cfg,
		db: conn,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
