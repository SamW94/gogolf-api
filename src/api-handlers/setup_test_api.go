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
		tx.Rollback()
		db.Close()
	}

	apiCfg := ApiConfig{
		DatabaseQueries: queries,
	}

	return apiCfg, cleanup
}
