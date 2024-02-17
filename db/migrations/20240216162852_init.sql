-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id VARCHAR(36),
  email VARCHAR(255),
  verified BOOLEAN DEFAULT FALSE,
);

CREATE TABLE links (
  id VARCHAR(10) PRIMARY KEY,
  name VARCHAR(50) NOT NULL DEFAULT 'unnamed',
  user_id VARCHAR(36) NOT NULL REFERENCES users(id),
  destination TEXT NOT NULL,
  created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE link_histories (
  id SERIAL PRIMARY KEY,
  user_agent TEXT,
  link_id VARCHAR(10) NOT NULL REFERENCES links(id),
  time_stamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
)

CREATE TABLE db_migrations (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TABLE db_migrations;
-- +goose StatementEnd
