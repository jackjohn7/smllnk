package users

import (
	"errors"

	"github.com/google/uuid"
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
	newUser := models.User{
		Id:       uuid.NewString(),
		Email:    email,
		Verified: false,
	}
	_, err = r.db.NamedExec(
		"INSERT INTO users (id, email, verified) VALUES(:id, :email, :verified)",
		newUser,
	)
	if err != nil {
		return nil, err
	}
	return &newUser, err
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

func (r *UserRepositoryPG) GetById(id string) (*models.User, error) {
	rows, err := r.db.Queryx("SELECT * FROM users u WHERE u.id = $1", id)
	if err != nil {
		return nil, err
	}
	users := make([]models.User, 0)
	for rows.Next() {
		var u models.User
		err = rows.StructScan(&u)
		users = append(users, u)
	}

	if len(users) != 1 {
		return nil, errors.New("No user found by this id")
	}

	return &users[0], nil
}

func (r *UserRepositoryPG) GetByEmail(email string) (*models.User, error) {
	rows, err := r.db.Queryx("SELECT * FROM users u WHERE u.email = $1", email)
	if err != nil {
		return nil, err
	}
	users := make([]models.User, 0)
	for rows.Next() {
		var u models.User
		err = rows.StructScan(&u)
		users = append(users, u)
	}

	if len(users) != 1 {
		return nil, errors.New("No user found by this email")
	}

	return &users[0], nil
}
