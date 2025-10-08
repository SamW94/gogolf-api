package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/SamW94/gogolf-api/src/internal/database"
	apiHandlers "github.com/SamW94/gogolf-api/src/api-handlers"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Printf("Failed to load .env file")
	}
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	serverPort := os.Getenv("INTERNAL_PORT")
	mux := http.NewServeMux()
	apiCfg := apiHandlers.ApiConfig{
		DatabaseQueries: dbQueries,
		Platform:        os.Getenv("PLATFORM"),
		JWTSecret:       os.Getenv("JWT_SECRET"),
		PolkaKey:        os.Getenv("POLKA_KEY"),
	}

	mux.HandleFunc("POST /api/golfers", apiCfg.CreateGolferHandler)

	server := &http.Server{
		Addr: ":" + serverPort,
		Handler: mux,
	}

	log.Printf("Serving on internal port: %s", serverPort)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
