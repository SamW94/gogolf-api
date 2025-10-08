-- name: CreateGolfer :one
INSERT INTO golfers (id, created_at, updated_at, email_address, username, hashed_password)
VALUES (gen_random_uuid (), NOW(), NOW(), $1, $2, $3)
RETURNING *;

-- name: RetrieveGolferByUsername :one
SELECT * from golfers
WHERE username = $1;