module github.com/SamW94/gogolf-api/src

go 1.25.2

replace github.com/SamW94/gogolf-api/src/internal/database => ./internal/database

replace github.com/SamW94/gogolf-api/src/internal/auth => ./internal/auth

replace github.com/SamW94/gogolf-api/src/api-handlers => ./api-handlers

require (
	github.com/SamW94/gogolf-api/src/api-handlers v0.0.0
	github.com/SamW94/gogolf-api/src/internal/database v0.0.0
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.10.9
)

require (
	github.com/SamW94/gogolf-api/src/internal/auth v0.0.0 // indirect
	github.com/golang-jwt/jwt/v5 v5.3.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	golang.org/x/crypto v0.43.0 // indirect
)
