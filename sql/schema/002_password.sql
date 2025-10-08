-- +goose Up
ALTER TABLE golfers
ADD COLUMN hashed_password TEXT NOT NULL
DEFAULT 'unset';

-- +goose Down
ALTER TABLE golfers
DROP COLUMN hashed_password;