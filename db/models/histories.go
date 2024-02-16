package models

import "time"

type LinkHistory struct {
	Id        int
	UserAgent string
	LinkId    string
	Timestamp time.Time
}
