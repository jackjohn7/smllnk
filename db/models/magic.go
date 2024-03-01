package models

import "time"

type MagicRequest struct {
	Id          string    `db:"id"`
	UserId      string    `db:"user_id"`
	CreatedDate time.Time `db:"created_date"`
	ExpiresAt   time.Time `db:"expires_at"`
}
