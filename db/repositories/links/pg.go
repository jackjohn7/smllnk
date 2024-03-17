package links

import (
	"errors"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/jackjohn7/smllnk/db/models"
	"github.com/jackjohn7/smllnk/utils"
)

type LinkRepositoryPG struct {
	db *sqlx.DB
}

func NewPG(db *sqlx.DB) *LinkRepositoryPG {
	return &LinkRepositoryPG{
		db: db,
	}
}

func (r *LinkRepositoryPG) Create(destination string, name string, creator *models.User) (link *models.Link, err error) {
	newLink := models.Link{
		Id:          utils.RandomString(7),
		Name:        name,
		Destination: destination,
		UserId:      creator.Id,
		CreatedDate: time.Now(),
	}

	_, err = r.db.NamedExec(
		"INSERT INTO links (id, name, user_id, destination, created_date) VALUES(:id, :name, :user_id, :destination, :created_date)",
		newLink,
	)
	if err != nil {
		return nil, err
	}
	return &newLink, nil
}

func (r *LinkRepositoryPG) Update(id string, newLink *models.Link) (ok bool) {
	return
}

// Delete link by id
func (r *LinkRepositoryPG) Delete(id string) bool {
	_, err := r.db.Exec("DELETE FROM links WHERE id=$1", id)
	return err == nil
}

// Get all links
func (r *LinkRepositoryPG) GetAll() (links []*models.Link, err error) {
	return nil, errors.New("Unimplemented")
}

// Get links by User
func (r *LinkRepositoryPG) GetAllUserLinks(user *models.User) (links []models.Link, err error) {
	rows, err := r.db.Queryx("SELECT * FROM links l WHERE l.user_id = $1", user.Id)
	if err != nil {
		return []models.Link{}, err
	}
	links = make([]models.Link, 0)
	for rows.Next() {
		var l models.Link
		err = rows.StructScan(&l)
		links = append(links, l)
	}

	return
}

// Delete all of a user's links
func (r *LinkRepositoryPG) DeleteAllUserLinks(user *models.User) (ok bool) {
	return
}

// Get link by Id
func (r *LinkRepositoryPG) GetById(id string) (link *models.Link, err error) {
	rows, err := r.db.Queryx("SELECT * FROM links l WHERE l.id = $1", id)
	if err != nil {
		return nil, err
	}

	links := make([]models.Link, 0)
	for rows.Next() {
		var l models.Link
		err = rows.StructScan(&l)
		links = append(links, l)
	}
	link = &links[0]
	return
}
