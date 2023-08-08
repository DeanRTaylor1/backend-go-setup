package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	authentication "github.com/deanrtaylor1/backend-go/internal/Auth"
	server "github.com/deanrtaylor1/backend-go/internal/Server"
	"github.com/deanrtaylor1/backend-go/internal/config"
	"github.com/deanrtaylor1/backend-go/internal/constants"
	db "github.com/deanrtaylor1/backend-go/internal/db/sqlc"
	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)

func connectToDB(config config.EnvConfig) (*sql.DB, error) {
	var (
		conn *sql.DB
		err  error
	)

	for i := 0; i < config.RetryCount; i++ {
		conn, err = sql.Open("postgres", config.ConnectionString)
		if err != nil {
			log.Printf("Cannot connect to db: %v", err)
			fmt.Println("Retrying in 15 seconds...")
			time.Sleep(15 * time.Second)
			continue
		}

		err = conn.Ping()
		if err != nil {
			log.Printf("Cannot ping db: %v", err)
			fmt.Println("Retrying in 15 seconds...")
			time.Sleep(15 * time.Second)
			continue
		}

		fmt.Println("Successfully connected to postgres.")
		return conn, nil // Successful connection
	}

	return nil, fmt.Errorf("could not connect to DB after multiple attempts: %v", err)
}

func main() {
	config.LoadEnv()
	config := config.Env
	fmt.Println(constants.FgCyan + "---------------------------" + constants.Reset)
	fmt.Println(constants.FgCyan + "LISTENING ON PORT: " + config.Port + constants.Reset)
	fmt.Println(constants.FgCyan + "--------------------------" + constants.BgBlack)

	conn, err := connectToDB(config)
	if err != nil {
		log.Fatalf("Cannot connect to db: %v", err)
	}

	// Create an instance of Queries with the db instance
	store := db.NewStore(conn)
	router := chi.NewMux()
	s := server.NewServer(config, store, router)
	authenticator, err := authentication.New()
	if err != nil {
		panic("Unable to create authenticator")
	}
	s.RegisterMiddlewares(*authenticator, store)
	s.RegisterRoutes(router)

	s.Start()
}
