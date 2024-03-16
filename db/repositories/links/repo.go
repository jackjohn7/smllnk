package links

import "github.com/jackjohn7/smllnk/db/models"

type LinkRepository interface {
	// Insert new link into database
	Create(destination string, name string, creator *models.User) (link *models.Link, err error)
	// Update link with new data
	Update(id string, newLink *models.Link) (ok bool)
	// Delete link by id
	Delete(id string) bool
	// Get all links
	GetAll() (links []*models.Link, err error)
	// Get links by User
	GetAllUserLinks(user *models.User) (links []models.Link, err error)
	// Delete all of a user's links
	DeleteAllUserLinks(user *models.User) (ok bool)
	// Get link by Id
	GetById(id string) (link *models.Link, err error)
}
