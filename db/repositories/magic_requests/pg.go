package magic_requests

import (
	"errors"
	"fmt"
	"time"

	"github.com/jackjohn7/smllnk/db/models"
	"github.com/jackjohn7/smllnk/utils"
	"github.com/jmoiron/sqlx"
)

type MagicRepositoryPG struct {
	db *sqlx.DB
}

func NewPG(db *sqlx.DB) *MagicRepositoryPG {
	return &MagicRepositoryPG{
		db: db,
	}
}

func (r *MagicRepositoryPG) Create(userId string) (*models.MagicRequest, error) {
	magicId, err := utils.GenerateMagicLinkId()
	if err != nil {
		return nil, err
	}
	newMagicRequest := models.MagicRequest{
		Id:          magicId,
		UserId:      userId,
		CreatedDate: time.Now(),
		ExpiresAt:   time.Now().Add(time.Minute * 20),
	}

	query := `INSERT INTO magic_requests
	(id, user_id, created_date, expires_at)
	VALUES(:id, :user_id, :created_date, :expires_at)`

	_, err = r.db.NamedExec(
		query,
		newMagicRequest)
	return &newMagicRequest, err
}

func (r *MagicRepositoryPG) Get(id string) (*models.MagicRequest, error) {
	rows, err := r.db.Queryx(
		"SELECT * FROM magic_requests mr WHERE mr.id = $1",
		id,
	)
	if err != nil {
		return nil, err
	}

	requests := make([]models.MagicRequest, 0)
	for rows.Next() {
		var mr models.MagicRequest
		err = rows.StructScan(&mr)
		requests = append(requests, mr)
	}

	if len(requests) != 1 {
		return nil, errors.New(fmt.Sprintf("Could not find unique MagicRequest where id=%s", id))
	}

	return &requests[0], nil
}

func (r *MagicRepositoryPG) Delete(id string) (ok bool) {
	query := `DELETE FROM magic_requests mr 
	WHERE mr.id = $1`
	_, err := r.db.Exec(query, id)
	if err == nil {
		ok = true
	}
	return
}
