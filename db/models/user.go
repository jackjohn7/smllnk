package models

type User struct {
	Id       string `db:"id"`
	Email    string `db:"email"`
	Verified bool   `db:"verified"`
}
