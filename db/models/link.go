package models

import "time"

type Link struct {
	Id          string    `db:"id"`
	Name        string    `db:"name"`
	UserId      string    `db:"user_id"`
	Destination string    `db:"destination"`
	CreatedDate time.Time `db:"created_date"`
}
