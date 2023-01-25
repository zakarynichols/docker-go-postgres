package db

import (
	"database/sql"
	"fmt"
	"time"
)

type Config struct {
	Password string
	User     string
	Name     string
	Host     string
	SSLMode  string
}

func Open(config Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=%s", config.User, config.Password, config.Name, config.Host, config.SSLMode)
	return sql.Open("postgres", connStr)
}

func queryNow(db *sql.DB) (time.Time, error) {
	var currentTime time.Time

	err := db.QueryRow("SELECT NOW()").Scan(&currentTime)
	if err != nil {
		return time.Time{}, err
	}

	return currentTime, err
}
