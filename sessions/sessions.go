package sessions

import (
	"time"

	"github.com/jackjohn7/smllnk/db/models"
)

type (
	Session struct {
		Id          string    `redis:"id"`
		UserId      string    `redis:"user_id"`
		UserAgent   string    `redis:"user_agent"`
		CreatedDate time.Time `redis:"created_date"`
		UpdatedDate time.Time `redis:"updated_date"`
		ExpiresAt   time.Time `redis:"expires_at"`
	}

	SessionStore interface {
		Create(user *models.User, userAgent string) (session *Session, err error)
		Get(id string) (session *Session, err error)
		Delete(id string) (ok bool)
		Refresh(id string) (session *Session, err error)
	}
)
