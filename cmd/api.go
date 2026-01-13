package main

import (
	"log"
	"net/http"
	"time"

	"database/sql"

	repo "github.com/Didul-arch/ecom-go-belajar/internal/adapters/mysql/sqlc"
	"github.com/Didul-arch/ecom-go-belajar/internal/orders"
	"github.com/Didul-arch/ecom-go-belajar/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// mount
func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	// user -> handler -> service -> db

	// A good base middleware stack
	r.Use(middleware.RequestID) // important for rate limiting
	r.Use(middleware.RealIP)    // for rate limiting and analytics and tracing
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/checkhealth", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Reps!"))
	})
	// http.ListenAndServe(":3000", r)

	productService := products.NewService(repo.New(app.db))
	productHandler := products.NewHandler(productService)

	orderService := orders.NewService(repo.New(app.db), app.db)
	ordersHandler := orders.NewHandler(orderService)

	r.Get("/products", productHandler.ListProducts)
	r.Post("/orders", ordersHandler.PlaceOrder)

	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server has started at address %s", app.config.addr)
	return srv.ListenAndServe()
}

type application struct {
	config config
	// logger
	// db driver
	db *sql.DB
}

type config struct {
	addr string // 8080, 12 factor document
	db   dbConfig
}

type dbConfig struct {
	dsn string // username, password, dbname, etc.
}
