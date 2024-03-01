package magic_requests

import "github.com/jackjohn7/smllnk/db/models"

type MagicRequestsRepository interface {
	// Create new magic request
	Create(userId string) (*models.MagicRequest, error)
	// Get magic request by Id
	Get(id string) (*models.MagicRequest, error)
	// Delete magic request by Id
	Delete(id string) (ok bool)
}
