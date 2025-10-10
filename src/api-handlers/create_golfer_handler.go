package apiHandlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/SamW94/gogolf-api/auth"
	"github.com/SamW94/gogolf-api/database"
	"github.com/google/uuid"
)

type requestJSONCreateGolfer struct {
	Password        string `json:"password"`
	RequestedGolfer string `json:"email"`
	Username        string `json:"username"`
}

type CreateGolferResponseSuccessful struct {
	ID           uuid.UUID `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	EmailAddress string    `json:"email"`
	Username     string    `json:"username"`
}

func (acfg *ApiConfig) CreateGolferHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	requestJson := requestJSONCreateGolfer{}
	err := decoder.Decode(&requestJson)
	if err != nil {
		log.Printf("Error decoding request body to requestJSON: %v", err)
		respondWithError(w, 500, "Something went wrong.")
		return
	}

	if requestJson.Password == "" {
		log.Printf("No password provided in request body.")
		respondWithError(w, 400, "No password provided in request body.")
		return
	} else {
		hashedPassword, err := auth.HashPassword(requestJson.Password)
		if err != nil {
			log.Printf("Error hashing password: %v", err)
			respondWithError(w, 500, "Something went wrong.")
			return
		}

		createUserParams := database.CreateGolferParams{
			EmailAddress:   requestJson.RequestedGolfer,
			HashedPassword: hashedPassword,
			Username:       requestJson.Username,
		}

		dbGolfer, err := acfg.DatabaseQueries.CreateGolfer(context.Background(), createUserParams)
		if err != nil {
			if err.Error() == "pq: duplicate key value violates unique constraint \"golfers_email_address_key\"" {
				log.Printf("Tried to create a new golfer but golfer with this email address already exists: %v", err)
				respondWithError(w, 409, "Email address is associated with an existing golfer.")
			} else if err.Error() == "pq: duplicate key value violates unique constraint \"golfers_username_key\"" {
				log.Printf("Tried to create a new golfer but a golfer with this username already exists: %v", err)
				respondWithError(w, 409, "Username is already in use.")
			} else {
				log.Printf("Error calling database.CreateGolfer() function: %v", err)
				respondWithError(w, 500, "Something went wrong.")
				return
			}
		} else {
			log.Printf("Successfully created golfer with ID %v, email %v, and username %v", dbGolfer.ID, dbGolfer.EmailAddress, dbGolfer.Username)
			respBody := CreateGolferResponseSuccessful{
				ID:           dbGolfer.ID,
				CreatedAt:    dbGolfer.CreatedAt,
				UpdatedAt:    dbGolfer.UpdatedAt,
				EmailAddress: dbGolfer.EmailAddress,
				Username:     dbGolfer.Username,
			}

			respondWithJSON(w, 201, respBody)
		}
	}
}
