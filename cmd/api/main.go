package main

import (
	"log"

	"social/internal/db"
	"social/internal/env"
	"social/internal/store"

	"github.com/joho/godotenv"
)

func main() {
	// Try to load .env, but don't crash if it's missing
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}
	cfg := config{
		addr: env.GetString("ADDRESS", ":8080"),
		db: dbConfig{
			addr:               env.GetString("DB_ADDR", "postgres://user:password@localhost/social?sslmode=disable"),
			maxOpenConnections: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConnections: env.GetInt("DB_MAX_IDLE_CONNS", 15),
			maxIdleTime:        env.GetString("DB_MAX_IDLE_TIME", "15min"),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConnections,
		cfg.db.maxIdleConnections,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("Db connected.")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
