package sessions

import (
	"time"

	"github.com/jackjohn7/smllnk/db/models"
)

type (
	Session struct {
		Id          string
		UserId      string
		UserAgent   string
		CreatedDate time.Time
		UpdatedDate time.Time
	}

	SessionStore interface {
		Create(user *models.User, userAgent string) (session *Session, err error)
		Get(id string) (session *Session, err error)
		Delete(id string) (ok bool)
		Refresh(id string) (session *Session, err error)
	}
)
