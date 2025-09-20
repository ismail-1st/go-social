package db

import (
	"context"
	"database/sql"
	"social/internal/env"
	"time"
)

// not using the db struct from the config because the internal package shouldnt know about other packages
func New(
	addr string,
	maxOpenConnections int,
	maxIdleConnections int,
	maxIdleTime string,
) (*sql.DB, error) {
	db, err := sql.Open("postgres", addr)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxIdleConnections)
	db.SetMaxIdleConns(maxIdleConnections)
	db.SetConnMaxIdleTime(env.GetDuration(maxIdleTime, "15mins"))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
