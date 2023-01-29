package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-postgres-docker/postgresql"
	"go-postgres-docker/redisdb"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	_ "github.com/lib/pq"
)

// Server extends the stdlib http.Server with the app's required services.
type Server struct {
	*http.Server
	userService UserService
}

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

	// User service
	userService := postgresql.NewUserService(psql)

	// Redis
	redis := redisdb.Open()

	// Router
	mux := mux.NewRouter()
	mux.Handle("/", http.HandlerFunc(handleRoot))
	mux.Handle("/now", http.HandlerFunc(handleNow(psql)))
	mux.Handle("/cache", http.HandlerFunc(pingRedis(ctx, redis)))

	// Cors
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowedHeaders: []string{"*"},
		MaxAge:         86400,
	})

	// Validate the port
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Init the server struct
	server := Server{
		Server: &http.Server{
			Addr:    ":" + port,
			Handler: c.Handler(mux),
		},
		userService: userService,
	}

	// Serve TLS
	log.Printf("Starting server on port%s\n", server.Addr)
	log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	responseData := map[string]string{"message": "Hello, World!"}
	jsonData, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

type currentTimeResponse struct {
	CurrentTime string `json:"current_time"`
}

type NowQuerier interface {
	QueryNow() (time.Time, error)
}

func handleNow(n NowQuerier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		time, err := n.QueryNow()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		responseData := currentTimeResponse{time.String()}
		jsonData, err := json.Marshal(responseData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

// pingRedis emits the redis cache is up and running.
func pingRedis(ctx context.Context, p Pinger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status, err := p.Ping(ctx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(status)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{
			"status": "Redis is up",
		}
		json.NewEncoder(w).Encode(response)
	}
}

type Pinger interface {
	Ping(context.Context) (string, error)
}

type UserService interface {
	QueryUsers() ([]postgresql.User, error)
}
