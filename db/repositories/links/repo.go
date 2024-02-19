package links

import "github.com/jackjohn7/smllnk/db/models"

type LinkRepository interface {
	// Insert new link into database
	Create(destination string, creator *models.User) (err error, link *models.Link)
	// Update link with new data
	Update(id string, newLink *models.Link) (ok bool)
	// Delete link by id
	Delete(id string) (ok bool)
	// Get all links
	GetAll() (err error, links []*models.Link)
	// Get links by User
	GetAllUserLinks(user *models.User) (err error, links []*models.Link)
	// Delete all of a user's links
	DeleteAllUserLinks(user *models.User) (ok bool)
	// Get link by Id
	GetById(id string) (err error, link *models.Link)
}
