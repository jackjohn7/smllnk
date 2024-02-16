package models

import "time"

type Link struct {
	Id          string
	UserId      string
	Destination string
	CreatedDate time.Time
}
