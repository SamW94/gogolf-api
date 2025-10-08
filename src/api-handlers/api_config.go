package apiHandlers

import (
	"github.com/SamW94/gogolf-api/src/internal/database"
)

type ApiConfig struct {
	DatabaseQueries *database.Queries
	Platform        string
	JWTSecret       string
	PolkaKey        string
}