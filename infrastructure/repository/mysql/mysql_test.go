package mysql_test

import (
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var testDb *sqlx.DB

func TestMain(m *testing.M) {
	err := godotenv.Load("../../../.env.testing")
	if err != nil {
		panic("Failed to load env. Error: " + err.Error())
	}

	testDb = sqlx.MustOpen(
		"mysql",
		os.Getenv("DB_TESTING_URL"),
	)
	if err := testDb.Ping(); err != nil {
		panic("Failed to ping test DB. Error: " + err.Error())
	}

	os.Exit(m.Run())
}
