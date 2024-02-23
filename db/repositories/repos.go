package repositories

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/jackjohn7/smllnk/db/repositories/links"
	"github.com/jackjohn7/smllnk/db/repositories/users"
	"github.com/jackjohn7/smllnk/environment"
)

type (
	Repositories struct {
		Users users.UserRepository
		Links links.LinkRepository
	}
)

func NewPGRepositories() *Repositories {
	// create pg connection
	db, err := sqlx.Connect("postgres", environment.Env.DbEnv.DATABASE_URL)
	if err != nil {
		log.Fatalln(err)
	}

	return &Repositories{
		Users: users.NewPG(db),
		Links: links.NewPG(db),
	}
}
