package repositories

import (
	"github.com/jmoiron/sqlx"

	"github.com/jackjohn7/smllnk/db/repositories/links"
	"github.com/jackjohn7/smllnk/db/repositories/users"
)

type (
	Repositories struct {
		Users users.UserRepository
		Links links.LinkRepository
	}
)

func NewPGRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Users: users.NewPG(db),
		Links: links.NewPG(db),
	}
}
