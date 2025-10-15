package apiHandlers

import (
	"database/sql"
	"os"
	"testing"

	"github.com/SamW94/gogolf-api/database"
	"github.com/joho/godotenv"
)

func setupTestAPI(t *testing.T) (ApiConfig, func()) {
	t.Helper()
	_ = godotenv.Load("../../.env")

	db, err := sql.Open("postgres", os.Getenv("TEST_DB_URL"))
	if err != nil {
		t.Fatalf("failed to open test DB: %v", err)
	}

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("failed to begin tx: %v", err)
	}

	queries := database.New(tx)

	cleanup := func() {
		err := tx.Rollback()
		if err != nil {
			t.Fatalf("WARNING: failed to rollback queries to the test DB, there may be inconsistencies: %s", err)
		}
		err = db.Close()
		if err != nil {
			t.Fatalf("Failed to close the connection to the database: %s", err)
		}
	}

	apiCfg := ApiConfig{
		DatabaseQueries: queries,
	}

	return apiCfg, cleanup
}
