-- +goose Up
-- +goose StatementBegin

CREATE TRIGGER update_table_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER IF EXISTS update_table_users_updated_at ON users;
-- +goose StatementEnd