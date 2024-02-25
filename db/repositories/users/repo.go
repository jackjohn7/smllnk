package users

import "github.com/jackjohn7/smllnk/db/models"

type UserRepository interface {
	// Insert new user into database
	Create(email string) (user *models.User, err error)
	// Update user with new data
	Update(id string, newUser *models.User) (ok bool)
	// Delete user by Id
	Delete(id string) (ok bool)
	// Get all Users
	GetAll() (users []*models.User, err error)
	// Get user of matching Id
	GetById(id string) (user *models.User, err error)
	// Get user of matching email
	GetByEmail(email string) (user *models.User, err error)
}
