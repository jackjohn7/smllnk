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

func (r *UserRepositoryPG) Create(email string) (user *models.User, err error) {
	return nil, errors.New("Unimplemented")
}

func (r *UserRepositoryPG) Update(id string, newUser *models.User) (ok bool) {
	return
}

func (r *UserRepositoryPG) Delete(id string) (ok bool) {
	return
}

func (r *UserRepositoryPG) GetAll() (users []*models.User, err error) {
	return []*models.User{}, errors.New("Unimplemented")
}

func (r *UserRepositoryPG) GetById(id string) (user *models.User, err error) {
	return nil, errors.New("Unimplemented")
}

func (r *UserRepositoryPG) GetByEmail(email string) (user *models.User, err error) {
	return nil, errors.New("Unimplemented")
}
