package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/zakarynichols/parent-teacher-portal/http"
	"github.com/zakarynichols/parent-teacher-portal/postgresql"
	"github.com/zakarynichols/parent-teacher-portal/redisdb"

	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()

	// Postgres
	config := postgresql.Config{
		Password: os.Getenv("POSTGRES_PASSWORD"),
		User:     os.Getenv("POSTGRES_USER"),
		Name:     os.Getenv("POSTGRES_DB"),
		Host:     os.Getenv("POSTGRES_HOST"),
		SSLMode:  os.Getenv("POSTGRES_SSLMODE"),
	}

	// Open psql
	psql, err := postgresql.Open(config)
	if err != nil {
		log.Fatal(err)
	}
	defer psql.Close()

	schoolService := postgresql.NewSchoolService(psql)

	// Redis
	redis := redisdb.Open()
	pong, err := redis.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong)

	port := "3000"

	server := http.New(
		port,
		schoolService,
	)

	server.RegisterSchoolRoutes(ctx)

	log.Fatal(server.ListenTLS())
}
