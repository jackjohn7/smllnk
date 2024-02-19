package users

import (
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/jackjohn7/smllnk/db/models"
)

type UserRepositoryPG struct {
	db *sqlx.DB
}

func NewPG(db *sqlx.DB) *UserRepositoryPG {
	return &UserRepositoryPG{
		db: db,
	}
}

func (r *UserRepositoryPG) Create(email string) (err error, user *models.User) {
	return errors.New("Unimplemented"), nil
}

func (r *UserRepositoryPG) Update(id string, newUser *models.User) (ok bool) {
	return
}

func (r *UserRepositoryPG) Delete(id string) (ok bool) {
	return
}

func (r *UserRepositoryPG) GetAll() (err error, users []*models.User) {
	return errors.New("Unimplemented"), []*models.User{}
}

func (r *UserRepositoryPG) GetById(id string) (err error, user *models.User) {
	return errors.New("Unimplemented"), nil
}

func (r *UserRepositoryPG) GetByEmail(email string) (err error, user *models.User) {
	return errors.New("Unimplemented"), nil
}
