package users

import "github.com/jackjohn7/smllnk/db/models"

type UserRepository interface {
	// Insert new user into database
	Create(email string) (err error, user *models.User)
	// Update user with new data
	Update(id string, newUser *models.User) (ok bool)
	// Delete user by Id
	Delete(id string) (ok bool)
	// Get all Users
	GetAll() (err error, users []*models.User)
	// Get user of matching Id
	GetById(id string) (err error, user *models.User)
	// Get user of matching email
	GetByEmail(email string) (err error, user *models.User)
}
