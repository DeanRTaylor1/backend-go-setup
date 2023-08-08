package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	config "github.com/deanrtaylor1/backend-go/internal/config"
	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config.LoadEnv()
	fmt.Printf("host: %s string: %s", config.Env.DbDriver, config.Env.ConnectionString)
	testDB, err := sql.Open(config.Env.DbDriver, config.Env.ConnectionString)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
