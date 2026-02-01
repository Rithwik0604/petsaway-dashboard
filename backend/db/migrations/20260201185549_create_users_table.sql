-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    users (
        id TEXT PRIMARY KEY,
        username TEXT NOT NULL UNIQUE,
        password_hash TEXT NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );

CREATE TRIGGER update_user_timestamp AFTER
UPDATE ON users BEGIN
UPDATE users
SET
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = NEW.id;

END;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE users;

-- +goose StatementEnd