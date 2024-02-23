-- +goose Up
-- +goose StatementBegin
DROP TABLE db_migrations;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
CREATE TABLE db_migrations (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL
);
-- +goose StatementEnd
