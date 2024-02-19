package repositories

import (
	"github.com/jackjohn7/smllnk/db/repositories/users"
	"github.com/jmoiron/sqlx"
)

type (
	Repository struct {
		Users users.UserRepositories
	}
)

func NewPGRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Users: users.NewPG(db),
	}
}
