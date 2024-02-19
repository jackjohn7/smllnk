package links

import (
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/jackjohn7/smllnk/db/models"
)

type LinkRepositoryPG struct {
	db *sqlx.DB
}

func NewPG(db *sqlx.DB) *LinkRepositoryPG {
	return &LinkRepositoryPG{
		db: db,
	}
}

func (r *LinkRepositoryPG) Create(destination string, creator *models.User) (err error, link *models.Link) {
	return errors.New("Unimplemented"), nil
}

func (r *LinkRepositoryPG) Update(id string, newLink *models.Link) (ok bool) {
	return
}

// Delete link by id
func (r *LinkRepositoryPG) Delete(id string) (ok bool) {
	return
}

// Get all links
func (r *LinkRepositoryPG) GetAll() (err error, links []*models.Link) {
	return errors.New("Unimplemented"), nil
}

// Get links by User
func (r *LinkRepositoryPG) GetAllUserLinks(user *models.User) (err error, links []*models.Link) {
	return errors.New("Unimplemented"), []*models.Link{}
}

// Delete all of a user's links
func (r *LinkRepositoryPG) DeleteAllUserLinks(user *models.User) (ok bool) {
	return
}

// Get link by Id
func (r *LinkRepositoryPG) GetById(id string) (err error, link *models.Link) {
	return errors.New("Unimplemented"), nil
}
