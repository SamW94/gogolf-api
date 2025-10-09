module github.com/SamW94/gogolf-api/src/api-handlers

go 1.25.2

replace github.com/SamW94/gogolf-api/src/auth => ../auth

replace github.com/SamW94/gogolf-api/src/database => ../database

require (
	github.com/SamW94/gogolf-api/src/auth v0.0.0-00010101000000-000000000000
	github.com/SamW94/gogolf-api/src/database v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.6.0
)

require (
	github.com/golang-jwt/jwt/v5 v5.3.0 // indirect
	golang.org/x/crypto v0.43.0 // indirect
)
