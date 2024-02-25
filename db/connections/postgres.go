package connections

import (
	"github.com/jackjohn7/smllnk/environment"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func NewPostgresConnection() (*sqlx.DB, error) {
	if db == nil {
		// create pg connection
		newDb, err := sqlx.Connect("postgres", environment.Env.DbEnv.DATABASE_URL)
		if err != nil {
			return nil, err
		}
		db = newDb
	}
	return db, nil
}
