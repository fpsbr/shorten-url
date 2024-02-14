package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"url-shortener/internal/data"

	_ "github.com/mattn/go-sqlite3"
)

const version = "1.0.0"

type config struct {
	port    int
	env     string
	baseUrl string
}

type application struct {
	config    config
	logger    *log.Logger
	shortener *data.URLShortener
}

func initDb(logger *log.Logger) *sql.DB {
	db, err := sql.Open("sqlite3", "url.db")
	if err != nil {
		logger.Fatal(err)
	}

	_, err = db.Exec(`
				CREATE TABLE IF NOT EXISTS url (
				    id INTEGER PRIMARY KEY,
				    url_code TEXT, 
				    short_url TEXT,
				    long_url TEXT,
				    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
				)
	`)
	if err != nil {
		logger.Fatal(err)
	}

	return db
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parsed()

	cfg.baseUrl = "http://localhost:8080"

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// db
	db := initDb(logger)
	//defer db.Close()

	app := &application{
		config:    cfg,
		logger:    logger,
		shortener: data.NewURLShortener(db, cfg.baseUrl),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
