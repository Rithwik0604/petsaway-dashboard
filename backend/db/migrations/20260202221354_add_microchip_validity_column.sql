-- +goose Up
-- +goose StatementBegin
ALTER TABLE clients ADD microchip_validity DATE;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE clients
DROP COLUMN microchip_validity;
-- +goose StatementEnd