package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/zakarynichols/parent-teacher-portal/http"
	"github.com/zakarynichols/parent-teacher-portal/postgres"
	"github.com/zakarynichols/parent-teacher-portal/redis"

	_ "github.com/lib/pq"
)

func main() {
	ctx := context.TODO()
	// Postgres
	pgConfig := postgres.Config{
		Password: os.Getenv("POSTGRES_PASSWORD"),
		User:     os.Getenv("POSTGRES_USER"),
		Name:     os.Getenv("POSTGRES_DB"),
		Host:     os.Getenv("POSTGRES_HOST"),
		SSLMode:  os.Getenv("POSTGRES_SSLMODE"),
	}

	// Open psql
	psql, err := postgres.Open(pgConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer psql.Close()

	schoolService := postgres.NewSchoolService(psql)

	// Redis
	redis := redis.Open()
	pong, err := redis.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong)

	port := "3000"

	httpConfig := http.Config{
		Addr:          port,
		SchoolService: schoolService,
	}

	server := http.New(httpConfig)

	server.RegisterSchoolRoutes(ctx)

	env := os.Getenv("APP_ENV")

	if env == "development" {
		log.Fatal(server.Listen())
	} else if env == "production" {
		log.Fatal(server.ListenTLS())
	} else {
		log.Fatal("ptp: malformed APP_ENV\n")
	}
}
