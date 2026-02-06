-- +goose Up
-- +goose StatementBegin
ALTER TABLE clients
ADD COLUMN notification_rabies INTEGER NOT NULL DEFAULT 0;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE clients
DROP COLUMN notification_rabies;

-- +goose StatementEnd