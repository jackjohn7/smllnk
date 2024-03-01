-- +goose Up
-- +goose StatementBegin
CREATE TABLE magic_requests (
  id VARCHAR(32) PRIMARY KEY,
  user_id VARCHAR(36) NOT NULL REFERENCES users(id),
  created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  expires_at TIMESTAMP NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE magic_requests;
-- +goose StatementEnd
