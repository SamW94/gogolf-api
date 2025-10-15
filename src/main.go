package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	apiHandlers "github.com/SamW94/gogolf-api/api-handlers"
	"github.com/SamW94/gogolf-api/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	_ = godotenv.Load("../.env")

	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	apiCfg := apiHandlers.ApiConfig{
		DatabaseQueries: dbQueries,
	}

	serverPort := os.Getenv("INTERNAL_PORT")
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/golfers", apiCfg.CreateGolferHandler)

	server := &http.Server{
		Addr:    ":" + serverPort,
		Handler: mux,
	}

	log.Printf("Serving on internal port: %s", serverPort)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
