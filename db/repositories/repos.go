package repositories

import (
	"log"

	_ "github.com/lib/pq"

	"github.com/jackjohn7/smllnk/db/connections"
	"github.com/jackjohn7/smllnk/db/repositories/links"
	"github.com/jackjohn7/smllnk/db/repositories/users"
)

type (
	Repositories struct {
		Users users.UserRepository
		Links links.LinkRepository
	}
)

func NewPGRepositories() *Repositories {
	// create pg connection
	db, err := connections.NewPostgresConnection()
	if err != nil {
		log.Fatalln(err)
	}

	return &Repositories{
		Users: users.NewPG(db),
		Links: links.NewPG(db),
	}
}
